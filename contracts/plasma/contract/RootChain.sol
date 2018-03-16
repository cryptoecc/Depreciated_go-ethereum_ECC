pragma solidity 0.4.18;


/**
 * @title Bytes operations
 *
 * @dev Based on https://github.com/GNSPS/solidity-bytes-utils/blob/master/contracts/BytesLib.sol
 */

library ByteUtils {
    function slice(bytes _bytes, uint _start, uint _length)
        internal
        pure
        returns (bytes)
    {

        bytes memory tempBytes;

        assembly {
            tempBytes := mload(0x40)

            let lengthmod := and(_length, 31)

            let mc := add(tempBytes, lengthmod)
            let end := add(mc, _length)

            for {
                let cc := add(add(_bytes, lengthmod), _start)
            } lt(mc, end) {
                mc := add(mc, 0x20)
                cc := add(cc, 0x20)
            } {
                mstore(mc, mload(cc))
            }

            mstore(tempBytes, _length)

            //update free-memory pointer
            //allocating the array padded to 32 bytes like the compiler does now
            mstore(0x40, and(add(mc, 31), not(31)))
        }

        return tempBytes;
    }
}

/**
 * @title Eliptic curve signature operations
 *
 * @dev Based on https://gist.github.com/axic/5b33912c6f61ae6fd96d6c4a47afde6d
 */

library ECRecovery {

  /**
   * @dev Recover signer address from a message by using his signature
   * @param hash bytes32 message, the hash is the signed message. What is recovered is the signer address.
   * @param sig bytes signature, the signature is generated using web3.eth.sign()
   */
    function recover(bytes32 hash, bytes sig)
        internal
        pure
        returns (address)
    {
        bytes32 r;
        bytes32 s;
        uint8 v;

        //Check the signature length
        if (sig.length != 65) {
        return (address(0));
        }

        // Divide the signature in v, r, and s variables
        assembly {
        r := mload(add(sig, 32))
        s := mload(add(sig, 64))
        v := byte(0, mload(add(sig, 96)))
        }

        // Version of signature should be 27 or 28, but 0 and 1 are also possible versions
        if (v < 27) {
        v += 27;
        }

        // If the version is correct return the signer address
        if (v != 27 && v != 28) {
        return (address(0));
        } else {
        return ecrecover(hash, v, r, s);
        }
    }
}


library Math {
    function max(uint256 a, uint256 b)
        internal
        pure
        returns (uint256)
    {
        if (a > b)
            return a;
        return b;
    }
}

library Merkle {
    function checkMembership(bytes32 leaf, uint256 index, bytes32 rootHash, bytes proof)
        internal
        pure
        returns (bool)
    {
        require(proof.length == 512);
        bytes32 proofElement;
        bytes32 computedHash = leaf;

        for (uint256 i = 32; i <= 512; i += 32) {
            assembly {
                proofElement := mload(add(proof, i))
            }
            if (index % 2 == 0) {
                computedHash = keccak256(computedHash, proofElement);
            } else {
                computedHash = keccak256(proofElement, computedHash);
            }
            index = index / 2;
        }
        return computedHash == rootHash;
    }
}

/**
* @title RLPReader
*
* RLPReader is used to read and parse RLP encoded data in memory.
*
* @author Andreas Olofsson (androlo1980@gmail.com)
*/
library RLP {

 uint constant DATA_SHORT_START = 0x80;
 uint constant DATA_LONG_START = 0xB8;
 uint constant LIST_SHORT_START = 0xC0;
 uint constant LIST_LONG_START = 0xF8;

 uint constant DATA_LONG_OFFSET = 0xB7;
 uint constant LIST_LONG_OFFSET = 0xF7;


 struct RLPItem {
     uint _unsafe_memPtr;    // Pointer to the RLP-encoded bytes.
     uint _unsafe_length;    // Number of bytes. This is the full length of the string.
 }

 struct Iterator {
     RLPItem _unsafe_item;   // Item that's being iterated over.
     uint _unsafe_nextPtr;   // Position of the next item in the list.
 }

 /* Iterator */

 function next(Iterator memory self) internal constant returns (RLPItem memory subItem) {
     if(hasNext(self)) {
         var ptr = self._unsafe_nextPtr;
         var itemLength = _itemLength(ptr);
         subItem._unsafe_memPtr = ptr;
         subItem._unsafe_length = itemLength;
         self._unsafe_nextPtr = ptr + itemLength;
     }
     else
         throw;
 }

 function next(Iterator memory self, bool strict) internal constant returns (RLPItem memory subItem) {
     subItem = next(self);
     if(strict && !_validate(subItem))
         throw;
     return;
 }

 function hasNext(Iterator memory self) internal constant returns (bool) {
     var item = self._unsafe_item;
     return self._unsafe_nextPtr < item._unsafe_memPtr + item._unsafe_length;
 }

 /* RLPItem */

 /// @dev Creates an RLPItem from an array of RLP encoded bytes.
 /// @param self The RLP encoded bytes.
 /// @return An RLPItem
 function toRLPItem(bytes memory self) internal constant returns (RLPItem memory) {
     uint len = self.length;
     if (len == 0) {
         return RLPItem(0, 0);
     }
     uint memPtr;
     assembly {
         memPtr := add(self, 0x20)
     }
     return RLPItem(memPtr, len);
 }

 /// @dev Creates an RLPItem from an array of RLP encoded bytes.
 /// @param self The RLP encoded bytes.
 /// @param strict Will throw if the data is not RLP encoded.
 /// @return An RLPItem
 function toRLPItem(bytes memory self, bool strict) internal constant returns (RLPItem memory) {
     var item = toRLPItem(self);
     if(strict) {
         uint len = self.length;
         if(_payloadOffset(item) > len)
             throw;
         if(_itemLength(item._unsafe_memPtr) != len)
             throw;
         if(!_validate(item))
             throw;
     }
     return item;
 }

 /// @dev Check if the RLP item is null.
 /// @param self The RLP item.
 /// @return 'true' if the item is null.
 function isNull(RLPItem memory self) internal constant returns (bool ret) {
     return self._unsafe_length == 0;
 }

 /// @dev Check if the RLP item is a list.
 /// @param self The RLP item.
 /// @return 'true' if the item is a list.
 function isList(RLPItem memory self) internal constant returns (bool ret) {
     if (self._unsafe_length == 0)
         return false;
     uint memPtr = self._unsafe_memPtr;
     assembly {
         ret := iszero(lt(byte(0, mload(memPtr)), 0xC0))
     }
 }

 /// @dev Check if the RLP item is data.
 /// @param self The RLP item.
 /// @return 'true' if the item is data.
 function isData(RLPItem memory self) internal constant returns (bool ret) {
     if (self._unsafe_length == 0)
         return false;
     uint memPtr = self._unsafe_memPtr;
     assembly {
         ret := lt(byte(0, mload(memPtr)), 0xC0)
     }
 }

 /// @dev Check if the RLP item is empty (string or list).
 /// @param self The RLP item.closeall
 /// @return 'true' if the item is null.
 function isEmpty(RLPItem memory self) internal constant returns (bool ret) {
     if(isNull(self))
         return false;
     uint b0;
     uint memPtr = self._unsafe_memPtr;
     assembly {
         b0 := byte(0, mload(memPtr))
     }
     return (b0 == DATA_SHORT_START || b0 == LIST_SHORT_START);
 }

 /// @dev Get the number of items in an RLP encoded list.
 /// @param self The RLP item.
 /// @return The number of items.
 function items(RLPItem memory self) internal constant returns (uint) {
     if (!isList(self))
         return 0;
     uint b0;
     uint memPtr = self._unsafe_memPtr;
     assembly {
         b0 := byte(0, mload(memPtr))
     }
     uint pos = memPtr + _payloadOffset(self);
     uint last = memPtr + self._unsafe_length - 1;
     uint itms;
     while(pos <= last) {
         pos += _itemLength(pos);
         itms++;
     }
     return itms;
 }

 /// @dev Create an iterator.
 /// @param self The RLP item.
 /// @return An 'Iterator' over the item.
 function iterator(RLPItem memory self) internal constant returns (Iterator memory it) {
     if (!isList(self))
         throw;
     uint ptr = self._unsafe_memPtr + _payloadOffset(self);
     it._unsafe_item = self;
     it._unsafe_nextPtr = ptr;
 }

 /// @dev Return the RLP encoded bytes.
 /// @param self The RLPItem.
 /// @return The bytes.
 function toBytes(RLPItem memory self) internal constant returns (bytes memory bts) {
     var len = self._unsafe_length;
     if (len == 0)
         return;
     bts = new bytes(len);
     _copyToBytes(self._unsafe_memPtr, bts, len);
 }

 /// @dev Decode an RLPItem into bytes. This will not work if the
 /// RLPItem is a list.
 /// @param self The RLPItem.
 /// @return The decoded string.
 function toData(RLPItem memory self) internal constant returns (bytes memory bts) {
     if(!isData(self))
         throw;
     var (rStartPos, len) = _decode(self);
     bts = new bytes(len);
     _copyToBytes(rStartPos, bts, len);
 }

 /// @dev Get the list of sub-items from an RLP encoded list.
 /// Warning: This requires passing in the number of items.
 /// @param self The RLP item.
 /// @return Array of RLPItems.
 function toList(RLPItem memory self, uint256 numItems) internal constant returns (RLPItem[] memory list) {
     if(!isList(self))
         throw;
     list = new RLPItem[](numItems);
     var it = iterator(self);
     uint idx;
     while(hasNext(it)) {
         list[idx] = next(it);
         idx++;
     }
 }

 /// @dev Decode an RLPItem into an ascii string. This will not work if the
 /// RLPItem is a list.
 /// @param self The RLPItem.
 /// @return The decoded string.
 function toAscii(RLPItem memory self) internal constant returns (string memory str) {
     if(!isData(self))
         throw;
     var (rStartPos, len) = _decode(self);
     bytes memory bts = new bytes(len);
     _copyToBytes(rStartPos, bts, len);
     str = string(bts);
 }

 /// @dev Decode an RLPItem into a uint. This will not work if the
 /// RLPItem is a list.
 /// @param self The RLPItem.
 /// @return The decoded string.
 function toUint(RLPItem memory self) internal constant returns (uint data) {
     if(!isData(self))
         throw;
     var (rStartPos, len) = _decode(self);
     if (len > 32)
         throw;
     assembly {
         data := div(mload(rStartPos), exp(256, sub(32, len)))
     }
 }

 /// @dev Decode an RLPItem into a boolean. This will not work if the
 /// RLPItem is a list.
 /// @param self The RLPItem.
 /// @return The decoded string.
 function toBool(RLPItem memory self) internal constant returns (bool data) {
     if(!isData(self))
         throw;
     var (rStartPos, len) = _decode(self);
     if (len != 1)
         throw;
     uint temp;
     assembly {
         temp := byte(0, mload(rStartPos))
     }
     if (temp > 1)
         throw;
     return temp == 1 ? true : false;
 }

    /// @dev Decode an RLPItem into a byte. This will not work if the
    /// RLPItem is a list.
    /// @param self The RLPItem.
    /// @return The decoded string.
    function toByte(RLPItem memory self)
        internal
        view
        returns (byte data)
    {
        require(isData(self));
        var (rStartPos, len) = _decode(self);
        if (len != 1)
            throw;
        uint temp;
        assembly {
            temp := byte(0, mload(rStartPos))
        }
        return byte(temp);
    }

    /// @dev Decode an RLPItem into an int. This will not work if the
    /// RLPItem is a list.
    /// @param self The RLPItem.
    /// @return The decoded string.
    function toInt(RLPItem memory self)
        internal
        view
        returns (int data)
    {
        return int(toUint(self));
    }

    /// @dev Decode an RLPItem into a bytes32. This will not work if the
    /// RLPItem is a list.
    /// @param self The RLPItem.
    /// @return The decoded string.
    function toBytes32(RLPItem memory self)
        internal
        view
        returns (bytes32 data)
    {
        return bytes32(toUint(self));
    }

    /// @dev Decode an RLPItem into an address. This will not work if the
    /// RLPItem is a list.
    /// @param self The RLPItem.
    /// @return The decoded string.
    function toAddress(RLPItem memory self)
        internal
        view
        returns (address data)
    {
        require(isData(self));
        var (rStartPos, len) = _decode(self);
        if (len != 20)
            throw;
        assembly {
            data := div(mload(rStartPos), exp(256, 12))
        }
    }

    // Get the payload offset.
    function _payloadOffset(RLPItem memory self)
        private
        view
        returns (uint)
    {
        if(self._unsafe_length == 0)
            return 0;
        uint b0;
        uint memPtr = self._unsafe_memPtr;
        assembly {
            b0 := byte(0, mload(memPtr))
        }
        if(b0 < DATA_SHORT_START)
            return 0;
        if(b0 < DATA_LONG_START || (b0 >= LIST_SHORT_START && b0 < LIST_LONG_START))
            return 1;
        if(b0 < LIST_SHORT_START)
            return b0 - DATA_LONG_OFFSET + 1;
        return b0 - LIST_LONG_OFFSET + 1;
    }

    // Get the full length of an RLP item.
    function _itemLength(uint memPtr)
        private
        view
        returns (uint len)
    {
        uint b0;
        assembly {
            b0 := byte(0, mload(memPtr))
        }
        if (b0 < DATA_SHORT_START)
            len = 1;
        else if (b0 < DATA_LONG_START)
            len = b0 - DATA_SHORT_START + 1;
        else if (b0 < LIST_SHORT_START) {
            assembly {
                let bLen := sub(b0, 0xB7) // bytes length (DATA_LONG_OFFSET)
                let dLen := div(mload(add(memPtr, 1)), exp(256, sub(32, bLen))) // data length
                len := add(1, add(bLen, dLen)) // total length
            }
        } else if (b0 < LIST_LONG_START) {
            len = b0 - LIST_SHORT_START + 1;
        } else {
            assembly {
                let bLen := sub(b0, 0xF7) // bytes length (LIST_LONG_OFFSET)
                let dLen := div(mload(add(memPtr, 1)), exp(256, sub(32, bLen))) // data length
                len := add(1, add(bLen, dLen)) // total length
            }
        }
    }

    // Get start position and length of the data.
    function _decode(RLPItem memory self)
        private
        view
        returns (uint memPtr, uint len)
    {
        require(isData(self));
        uint b0;
        uint start = self._unsafe_memPtr;
        assembly {
            b0 := byte(0, mload(start))
        }
        if (b0 < DATA_SHORT_START) {
            memPtr = start;
            len = 1;
            return;
        }
        if (b0 < DATA_LONG_START) {
            len = self._unsafe_length - 1;
            memPtr = start + 1;
        } else {
            uint bLen;
            assembly {
                bLen := sub(b0, 0xB7) // DATA_LONG_OFFSET
            }
            len = self._unsafe_length - 1 - bLen;
            memPtr = start + bLen + 1;
        }
        return;
    }

    // Assumes that enough memory has been allocated to store in target.
    function _copyToBytes(uint btsPtr, bytes memory tgt, uint btsLen)
        private
        view
    {
        // Exploiting the fact that 'tgt' was the last thing to be allocated,
        // we can write entire words, and just overwrite any excess.
        assembly {
            {
                    let i := 0 // Start at arr + 0x20
                    let words := div(add(btsLen, 31), 32)
                    let rOffset := btsPtr
                    let wOffset := add(tgt, 0x20)
                tag_loop:
                    jumpi(end, eq(i, words))
                    {
                        let offset := mul(i, 0x20)
                        mstore(add(wOffset, offset), mload(add(rOffset, offset)))
                        i := add(i, 1)
                    }
                    jump(tag_loop)
                end:
                    mstore(add(tgt, add(0x20, mload(tgt))), 0)
            }
        }
    }

    // Check that an RLP item is valid.
    function _validate(RLPItem memory self)
        private
        pure
        returns (bool ret)
    {
        // Check that RLP is well-formed.
        uint b0;
        uint b1;
        uint memPtr = self._unsafe_memPtr;
        assembly {
            b0 := byte(0, mload(memPtr))
            b1 := byte(1, mload(memPtr))
        }
        if(b0 == DATA_SHORT_START + 1 && b1 < DATA_SHORT_START)
            return false;
        return true;
    }
}


/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {
    function mul(uint256 a, uint256 b)
        internal
        pure
        returns (uint256)
    {
        if (a == 0) {
        return 0;
        }
        uint256 c = a * b;
        assert(c / a == b);
        return c;
    }

    function div(uint256 a, uint256 b)
        internal
        pure
        returns (uint256)
    {
        // assert(b > 0); // Solidity automatically throws when dividing by 0
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold
        return c;
    }

    function sub(uint256 a, uint256 b)
        internal
        pure
        returns (uint256)
    {
        assert(b <= a);
        return a - b;
    }

    function add(uint256 a, uint256 b)
        internal
        pure
        returns (uint256)
    {
        uint256 c = a + b;
        assert(c >= a);
        return c;
    }
}


library Validate {
    function checkSigs(bytes32 txHash, bytes32 rootHash, uint256 inputCount, bytes sigs)
        internal
        view
        returns (bool)
    {
        require(sigs.length % 65 == 0 && sigs.length <= 260);
        bytes memory sig1 = ByteUtils.slice(sigs, 0, 65);
        bytes memory sig2 = ByteUtils.slice(sigs, 65, 65);
        bytes memory confSig1 = ByteUtils.slice(sigs, 130, 65);
        bytes32 confirmationHash = keccak256(txHash, rootHash);
        if (inputCount == 0) {
            return msg.sender == ECRecovery.recover(confirmationHash, confSig1);
        }
        if (inputCount < 1000000000) {
            return ECRecovery.recover(txHash, sig1) == ECRecovery.recover(confirmationHash, confSig1);
        } else {
            bytes memory confSig2 = ByteUtils.slice(sigs, 195, 65);
            bool check1 = ECRecovery.recover(txHash, sig1) == ECRecovery.recover(confirmationHash, confSig1);
            bool check2 = ECRecovery.recover(txHash, sig2) == ECRecovery.recover(confirmationHash, confSig2);
            return check1 && check2;
        }
    }
}


contract PriorityQueue {
    using SafeMath for uint256;

    /*
     *  Modifiers
     */
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    /*
     *  Storage
     */
    address owner;
    uint256[] heapList;
    uint256 public currentSize;

    function PriorityQueue()
        public
    {
        owner = msg.sender;
        heapList = [0];
        currentSize = 0;
    }

    function insert(uint256 k)
        public
        onlyOwner
    {
        heapList.push(k);
        currentSize = currentSize.add(1);
        percUp(currentSize);
    }

    function minChild(uint256 i)
        public
        view
        returns (uint256)
    {
        if (i.mul(2).add(1) > currentSize) {
            return i.mul(2);
        } else {
            if (heapList[i.mul(2)] < heapList[i.mul(2).add(1)]) {
                return i.mul(2);
            } else {
                return i.mul(2).add(1);
            }
        }
    }

    function getMin()
        public
        view
        returns (uint256)
    {
        return heapList[1];
    }

    function delMin()
        public
        onlyOwner
        returns (uint256)
    {
        uint256 retVal = heapList[1];
        heapList[1] = heapList[currentSize];
        delete heapList[currentSize];
        currentSize = currentSize.sub(1);
        percDown(1);
        return retVal;
    }

    function percUp(uint256 i)
        private
    {
        while (i.div(2) > 0) {
            if (heapList[i] < heapList[i.div(2)]) {
                uint256 tmp = heapList[i.div(2)];
                heapList[i.div(2)] = heapList[i];
                heapList[i] = tmp;
            }
            i = i.div(2);
        }
    }

    function percDown(uint256 i)
        private
    {
        while (i.mul(2) <= currentSize) {
            uint256 mc = minChild(i);
            if (heapList[i] > heapList[mc]) {
                uint256 tmp = heapList[i];
                heapList[i] = heapList[mc];
                heapList[mc] = tmp;
            }
            i = mc;
        }
    }
}
/**
 * @title RootChain
 * @dev This contract secures a utxo payments plasma child chain to ethereum
 */


contract RootChain {
    using SafeMath for uint256;
    using RLP for bytes;
    using RLP for RLP.RLPItem;
    using RLP for RLP.Iterator;
    using Merkle for bytes32;

    /*
     * Events
     */
    event Deposit(address depositor, uint256 amount);
    event Exit(address exitor, uint256 utxoPos);

    /*
     *  Storage
     */
    mapping(uint256 => childBlock) public childChain;
    mapping(uint256 => exit) public exits;
    mapping(uint256 => uint256) public exitIds;
    PriorityQueue exitsQueue;
    address public authority;
    uint256 public currentChildBlock;
    uint256 public recentBlock;
    uint256 public weekOldBlock;

    struct exit {
        address owner;
        uint256 amount;
        uint256 utxoPos;
    }

    struct childBlock {
        bytes32 root;
        uint256 created_at;
    }

    /*
     *  Modifiers
     */
    modifier isAuthority() {
        require(msg.sender == authority);
        _;
    }

    modifier incrementOldBlocks() {
        while (childChain[weekOldBlock].created_at < block.timestamp.sub(1 weeks)) {
            if (childChain[weekOldBlock].created_at == 0)
                break;
            weekOldBlock = weekOldBlock.add(1);
        }
        _;
    }

    function RootChain()
        public
    {
        authority = msg.sender;
        currentChildBlock = 1;
        exitsQueue = new PriorityQueue();
    }

    // @dev Allows Plasma chain operator to submit block root
    // @param root The root of a child chain block
    function submitBlock(bytes32 root, uint256 blknum)
        public
        isAuthority
        incrementOldBlocks
    {
        require(blknum == currentChildBlock);
        childChain[currentChildBlock] = childBlock({
            root: root,
            created_at: block.timestamp
        });
        currentChildBlock = currentChildBlock.add(1);
    }

    // @dev Allows anyone to deposit funds into the Plasma chain
    // @param txBytes The format of the transaction that'll become the deposit
    // TODO: This needs to be optimized so that the transaction is created
    //       from msg.sender and msg.value
    function deposit(bytes txBytes)
        public
        payable
    {
        var txList = txBytes.toRLPItem().toList(11);
        require(txList.length == 11);
        for (uint256 i; i < 6; i++) {
            require(txList[i].toUint() == 0);
        }
        require(txList[7].toUint() == msg.value);
        require(txList[9].toUint() == 0);
        bytes32 zeroBytes;
        bytes32 root = keccak256(keccak256(txBytes), new bytes(130));
        for (i = 0; i < 16; i++) {
            root = keccak256(root, zeroBytes);
            zeroBytes = keccak256(zeroBytes, zeroBytes);
        }
        childChain[currentChildBlock] = childBlock({
            root: root,
            created_at: block.timestamp
        });
        currentChildBlock = currentChildBlock.add(1);
        Deposit(txList[6].toAddress(), txList[7].toUint());
    }

    // @dev Starts to exit a specified utxo
    // @param utxoPos The position of the exiting utxo in the format of blknum * 1000000000 + index * 10000 + oindex
    // @param txBytes The transaction being exited in RLP bytes format
    // @param proof Proof of the exiting transactions inclusion for the block specified by utxoPos
    // @param sigs Both transaction signatures and confirmations signatures used to verify that the exiting transaction has been confirmed
    function startExit(uint256 utxoPos, bytes txBytes, bytes proof, bytes sigs)
        public
        incrementOldBlocks
    {
        var txList = txBytes.toRLPItem().toList(11);
        uint256 blknum = utxoPos / 1000000000;
        uint256 txindex = (utxoPos % 1000000000) / 10000;
        uint256 oindex = utxoPos - blknum * 1000000000 - txindex * 10000;
        bytes32 root = childChain[blknum].root;

        require(msg.sender == txList[6 + 2 * oindex].toAddress());
        bytes32 txHash = keccak256(txBytes);
        bytes32 merkleHash = keccak256(txHash, ByteUtils.slice(sigs, 0, 130));
        uint256 inputCount = txList[3].toUint() * 1000000000 + txList[0].toUint();
        require(Validate.checkSigs(txHash, root, inputCount, sigs));
        require(merkleHash.checkMembership(txindex, root, proof));

        // Priority is a given utxos position in the exit priority queue
        uint256 priority;
        if (blknum < weekOldBlock) {
            priority = (utxoPos / blknum).mul(weekOldBlock);
        } else {
            priority = utxoPos;
        }
        require(exitIds[utxoPos] == 0);
        exitIds[utxoPos] = priority;
        exitsQueue.insert(priority);
        exits[priority] = exit({
            owner: txList[6 + 2 * oindex].toAddress(),
            amount: txList[7 + 2 * oindex].toUint(),
            utxoPos: utxoPos
        });
        Exit(msg.sender, utxoPos);
    }

    // @dev Allows anyone to challenge an exiting transaction by submitting proof of a double spend on the child chain
    // @param cUtxoPos The position of the challenging utxo
    // @param eUtxoPos The position of the exiting utxo
    // @param txBytes The challenging transaction in bytes RLP form
    // @param proof Proof of inclusion for the transaction used to challenge
    // @param sigs Signatures for the transaction used to challenge
    // @param confirmationSig The confirmation signature for the transaction used to challenge
    function challengeExit(uint256 cUtxoPos, uint256 eUtxoPos, bytes txBytes, bytes proof, bytes sigs, bytes confirmationSig)
        public
    {
        uint256 txindex = (cUtxoPos % 1000000000) / 10000;
        bytes32 root = childChain[cUtxoPos / 1000000000].root;
        uint256 priority = exitIds[eUtxoPos];
        var txHash = keccak256(txBytes);
        var confirmationHash = keccak256(txHash, root);
        var merkleHash = keccak256(txHash, sigs);
        address owner = exits[priority].owner;

        require(owner == ECRecovery.recover(confirmationHash, confirmationSig));
        require(merkleHash.checkMembership(txindex, root, proof));
        delete exits[priority];
        delete exitIds[eUtxoPos];
    }


    // @dev Loops through the priority queue of exits, settling the ones whose challenge
    // @dev challenge period has ended
    function finalizeExits()
        public
        incrementOldBlocks
        returns (uint256)
    {
        uint256 twoWeekOldTimestamp = block.timestamp.sub(2 weeks);
        exit memory currentExit = exits[exitsQueue.getMin()];
        uint256 blknum = currentExit.utxoPos.div(1000000000);
        while (childChain[blknum].created_at < twoWeekOldTimestamp && exitsQueue.currentSize() > 0) {
            currentExit.owner.transfer(currentExit.amount);
            uint256 priority = exitsQueue.delMin();
            delete exits[priority];
            delete exitIds[currentExit.utxoPos];
            currentExit = exits[exitsQueue.getMin()];
        }
    }

    /*
     *  Constants
     */
    function getChildChain(uint256 blockNumber)
        public
        view
        returns (bytes32, uint256)
    {
        return (childChain[blockNumber].root, childChain[blockNumber].created_at);
    }

    function getExit(uint256 priority)
        public
        view
        returns (address, uint256, uint256)
    {
        return (exits[priority].owner, exits[priority].amount, exits[priority].utxoPos);
    }
}
