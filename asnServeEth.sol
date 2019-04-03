pragma solidity^0.4.25;

//import "./SafeMath.sol";
//import "./stringUtils.sol";

// https://github.com/ethereum/EIPs/issues/20
interface Token {
    function totalSupply() external view returns (uint supply);
    function balanceOf(address _owner) external view returns (uint balance);
    function transfer(address _to, uint _value) external returns (bool success);
    function transferFrom(address _from, address _to, uint _value) external returns (bool success);
    function approve(address _spender, uint _value) external returns (bool success);
    function allowance(address _owner, address _spender) external view returns (uint remaining);
    //function decimals() external view returns(uint digits);
    event Transfer(address indexed _from, address indexed _to, uint _value);
    event Approval(address indexed _owner, address indexed _spender, uint _value);
}

contract Utils {
    Token constant internal TOKEN_ETH_ADDRESS = Token(0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa);
    // uint constant internal ETH_DECIMALS = 18;
    // uint constant internal MAX_DECIMALS = 18;

    // uint constant internal MAX_PRICE_PRODUCT = 10 ** 16;
    // uint constant internal PRICE_PRODUCT = 10 ** 8;

    // uint constant internal MAX_VALUE = 2 ** 255 - 1;

    // mapping(address => uint) tokenDecimals;

    // function setTokenDecimals(Token _token) public {
    //     if (TOKEN_ETH_ADDRESS == _token) {
    //         tokenDecimals[_token] = ETH_DECIMALS;
    //     } else {
    //         tokenDecimals[_token] = _token.decimals();
    //     }
    // }

    // function getTokenDecimals(Token _token) public view returns(uint) {
    //     if (TOKEN_ETH_ADDRESS == _token) {
    //         return MAX_DECIMALS;
    //     } else {
    //         return _token.decimals();
    //     }
    // }

    // function getBalance(Token _token, address _user) public view returns(uint) {
    //     if(TOKEN_ETH_ADDRESS == _token) {
    //         return _user.balance;
    //     } else {
    //         return _token.balanceOf(_user);
    //     }
    // }
}

contract ASN is Utils{
    //using SafeMath for uint256;

    //using StringUtils for string;

    string public contractName;

    struct Transaction {
        uint64  id;
        bytes32 hashSecret;
        address sender;
        address receiver;
        Token token;
        uint amount;
        uint refundStartTime;            // refundStartTime = 2 * redeemEndTime
        uint redeemEndTime;          // redeemEndTime < refundStartTime, receiver's expiration is less than the sender's.
        address settler;
        string secret;
    }

    mapping(uint64 => Transaction) transactions;       // all transactions

    constructor() public {
        contractName = "ASN";
    }

    function() public {
        revert("The ASN contract can not reveive Ether directly!");
    }

    event Fund(uint64 _id, bytes32 _hashSecretX, address _sender, address _receiver, Token _token, uint _amount, uint _expirationSender, uint _expirationRecerver);

    function fund(uint64 _id, bytes32 _hashSecretX, address _receiver, Token _token, uint _amount, uint _expirationSender, uint _expirationReceiver)
        public
        payable
        returns(bool)
    {
        if(_token == TOKEN_ETH_ADDRESS) {
            require(_amount == msg.value, "Amount must equal to msg.value.");
        } else {
            require(msg.value == 0, "msg.value must equal to 0.");
            require(_token.transferFrom(msg.sender, this, _amount), "Sender sends token to ASN contract.");
        }

        uint refundStartTime = _expirationSender * 1;

        transactions[_id] = Transaction(_id, _hashSecretX, msg.sender, _receiver, _token, _amount, refundStartTime, _expirationReceiver, address(0x0), "");

        emit Fund(_id, _hashSecretX, msg.sender, _receiver, _token, _amount, refundStartTime, _expirationReceiver);

        return true;
    }

    event Refund(uint64 _id, address _sender, address _receiver, Token _token, uint _amount);

    function refund(uint64 _id) public returns(bool) {
        require(msg.sender == transactions[_id].sender, "Correct sender.");

        address sender = transactions[_id].sender;
        address receiver = transactions[_id].receiver;
        uint amount = transactions[_id].amount;
        Token token = transactions[_id].token;

        require(transactions[_id].settler == address(0x0), "Transaction has not been settled.");

        //now = block.timestamp  当前区块的时间戳
        require(block.timestamp >= transactions[_id].refundStartTime, "Sender's expiration");

        transactions[_id].settler = msg.sender;

        if(token == TOKEN_ETH_ADDRESS) {
            (sender).transfer(amount);
        } else {
            require(token.transfer(sender, amount), "Refund token.");
        }

        emit Refund(_id, sender, receiver, token, amount);

        return true;
    }

    event Redeem(uint64 _id, address _sender, address _receiver, Token _token, uint _amount, string _secret);

    function redeem(uint64 _id, string _secret) public returns(bool) {
        bytes32 hashSecret = sha256((_secret));

        //string memory hashSecretString = bytes32ToStr(hashSecret);

        require(msg.sender == transactions[_id].receiver, "Correct receiver.");
        require(hashSecret == transactions[_id].hashSecret, "Correct secret.");

        //require(StringUtils.equal(hashSecretString,transactions[_id].hashSecret),"Correct secret.");

        require(block.timestamp <= transactions[_id].redeemEndTime, "Receiver's expiration");
        //require(block.timestamp >= transactions[_id].redeemEndTime, "Receiver's expiration");

        require(transactions[_id].settler == address(0x0), "Transaction has not been settled.");

        address sender = transactions[_id].sender;
        address receiver = msg.sender;
        Token token = transactions[_id].token;
        uint amount = transactions[_id].amount;

        transactions[_id].settler = receiver;
        transactions[_id].secret = _secret;

        if(token == TOKEN_ETH_ADDRESS) {
            receiver.transfer(amount);
        } else {
            require(token.transfer(receiver, amount), "Receive fromToken.");
        }

        emit Redeem(_id, sender, receiver, token, amount, _secret);

        return true;
    }
    
    function redeemKeccak(string _secret) constant returns(bytes32, string) {
        bytes32 hashSecret = sha256((_secret));
        string memory hashSecretString = bytes32ToStr(hashSecret);
        
        return (hashSecret, hashSecretString);
    }

    // function check(uint64 _id, string hashSecret, address sender, address receiver, address token, uint amount, uint refundStartTime, uint redeemEndTime) public view returns(bool) {

    // }

    function getHashSecretByid(uint64 _id) public view returns(bytes32) {
        require(msg.sender == transactions[_id].sender || msg.sender == transactions[_id].receiver, "Msg.sender in the transaction.");
        return transactions[_id].hashSecret;
    }

    function getSecretByid(uint64 _id) public view returns(string) {
        require(msg.sender == transactions[_id].sender || msg.sender == transactions[_id].receiver, "Msg.sender in the transaction.");
        return transactions[_id].secret;
    }

    function getTransactionByid(uint64 _id)
        public
        view
        returns(bytes32 hashSecret, address sender, address receiver, address token, uint amount, uint refundStartTime, uint redeemEndTime, address settler, string secret)
    {
        require(msg.sender == transactions[_id].sender || msg.sender == transactions[_id].receiver, "Msg.sender in the transaction.");

        hashSecret = transactions[_id].hashSecret;
        sender = transactions[_id].sender;
        receiver = transactions[_id].receiver;
        token = transactions[_id].token;
        amount = transactions[_id].amount;
        refundStartTime = transactions[_id].refundStartTime;
        redeemEndTime = transactions[_id].redeemEndTime;
        settler = transactions[_id].settler;
        secret = transactions[_id].secret;
    }
    
    function bytes32ToStr(bytes32 _bytes32) public constant returns (string){
        bytes memory bytesArray = new bytes(32);
        for (uint256 i; i < 32; i++) {
            bytesArray[i] = _bytes32[i];
        }
        return string(bytesArray);
    }
    
}
