contract A {
    address public implementation;
    uint256 public x;

    function add() external returns (uint256) {
        return x = x + 1;
    }
}
