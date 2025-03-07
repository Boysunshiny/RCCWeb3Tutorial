// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Todo {
    struct TodoItem {
        string name;
        bool isCompleted;
    }
    TodoItem[] public s_todoList; // 29414

    function create(string memory name) public {
        s_todoList.push(TodoItem({name: name, isCompleted: false}));
    }

    function modifyName(uint index, string memory name) public {
        s_todoList[index].name = name;
    }

    function modiName2(uint256 index_, string memory name_) external {
        // 方法2: 先获取储存到 storage，在修改，在修改多个属性的时候比较省 gas
        TodoItem storage temp = s_todoList[index_];
        temp.name = name_;
    }

    function toggle(uint index) public {
        s_todoList[index].isCompleted = !s_todoList[index].isCompleted;
    }

    function get(
        uint index
    ) public view returns (string memory name, bool isCompleted) {
        TodoItem storage temp = s_todoList[index];
        return (temp.name, temp.isCompleted);
    }

    // 获取任务1: memory : 2次拷贝
    // 29448 gas
    function get1(
        uint256 index_
    ) external view returns (string memory name_, bool status_) {
        TodoItem memory temp = s_todoList[index_];
        return (temp.name, temp.isCompleted);
    }

    // 获取任务2: storage : 1次拷贝
    // 预期：get2 的 gas 费用比较低（相对 get1）
    // 29388 gas
    function get2(
        uint256 index_
    ) external view returns (string memory name_, bool status_) {
        TodoItem storage temp = s_todoList[index_];
        return (temp.name, temp.isCompleted);
    }
}
