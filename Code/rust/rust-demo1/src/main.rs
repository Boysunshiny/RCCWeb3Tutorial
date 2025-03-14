use std::arch::x86_64;
use std::collections::{HashMap, hash_map};
use std::fmt::{self, Error};
use std::fmt::{Display, Formatter, Result};
use std::fs::{self, File};
use std::io::{self, Read};
use std::ops::{Add, Index};
use std::rc::Rc;
use std::slice::from_raw_parts;
use std::str::from_utf8_unchecked;
use std::sync::Arc;
use std::{array, result};
use std::{iter, string, thread};

fn main() {
    let a = "healle";
    let mut b = String::from("value");

    // std::ops::Add::add(b, "323");
    let b = String::add(b, "rhs");
    let c = 3;
    let d = std::ops::Add::add(c, 3);
    let mut c = &b[..];
    println!("{a} {b} {c}");

    let mut b = vec![1, 2, 3, 4, 5, 6];

    let b = Box::new(a);
    // let x: [i32] =

    let mut b = [1, 23, 4, 5];
    let b = b.iter_mut();

    let b = [1, 23, 4, 5];
    let b = b.iter();

    let mut b = [1, 23, 4, 5];
    let b = b.into_iter();

    println!("{b:?}");
    // let d = String::from("value");
    // b = String::add(b, "323");
    // let s = b + d.as_str();
    // println!("{b} {s}");

    //String vec<T>
    //      [T;n]
    // str  [T]
    // &str &[T]

    let c = Cat {
        name: String::from("cat"),
    };

    c.bark();
    Say::bark(&c);
    <Cat as Say>::bark(&c);

    let (pointer, length) = get_memory_location();
    let message = get_str_at_location(pointer, length);
    println!(
        "The {} bytes at 0x{:X} stored: {}",
        length, pointer, message
    );
    // 如果大家想知道为何处理裸指针需要 `unsafe`，可以试着反注释以下代码
    let message = get_str_at_location(1000, 10);

    let r1;
    let r2;
    {
        static STATIC_EXAMPLE: i32 = 42;
        r1 = &STATIC_EXAMPLE;
        let x = "&'static str";
        r2 = x;
        // r1 和 r2 持有的数据都是 'static 的，因此在花括号结束后，并不会被释放
    }
    println!("&'static i32: {}", r1); // -> 42
    println!("&'static str: {}", r2); // -> &'static str
    // println!("{message}")\
    let b = 3.1 as i8;
    let c = 300i32;
    let d = c as i8;

    let mut b = [1, 2, 3, 4, 5, 6];

    let p = b.as_mut_ptr();
    let x = p as usize;
    let x = x + 4;
    let p = x as *mut i32;
    unsafe {
        *p = 100;
    }
    println!("{b:?}");
    let b = 300;

    let x: i8 = match b.try_into() {
        Ok(c) => c,
        Err(_) => 0,
    };

    println!("{x}");
    println!("22");

    let t: &mut i32 = &mut 0;
    foo2(t);

    let b = TestDeref {
        name: String::from("hello"),
        value: vec![(1, 121), (3, 121), (2, 2323)].into_iter().collect(),
    };
    b.test();

    let b = &TestDeref {
        name: String::from("hello"),
        value: vec![(1, 121), (3, 121), (2, 2323)].into_iter().collect(),
    };
    b.test();

    let b = &mut TestDeref {
        name: String::from("hello"),
        value: vec![(1, 121), (3, 121), (2, 2323)].into_iter().collect(),
    };
    b.test();
    // b.testGen::<i32>(&c);

    let b = b[1];
    let c = 1;

    let array: Rc<Box<[i32; 6]>> = Rc::new(Box::new([1, 2, 3, 4, 5, 6]));
    let t = array[0];
    let t = array.index(0);
    let t = Index::index(&array as &[i32; 6], 1);
    // Index::<i32>::index(&array, 0);
    // Index::<[i32; 6]>::index(&array, 0);
    // let t = Index::<i32>::index(&array, 1);

    let b = 12121;
    let c = do_stuff2(&b);

    println!("{c}");

    let b: *const () = foo as *const ();
    let func = unsafe { std::mem::transmute::<*const (), fn() -> i8>(b) };
    assert_eq!(func(), 2)
}

fn foo() -> i32 {
    2
}
fn do_stuff<T: Clone>(value: &T) {
    let cloned = value.clone();
}
fn do_stuff2<T>(value: &T) -> &T
where
    T: Add<Output = T>,
{
    let mut cloned = value.clone();
    cloned
}
#[derive(Clone)]
struct Container<T>(Arc<T>);

fn clone_containers<T>(foo: &Container<i32>, bar: &Container<T>) {
    let foo_cloned = foo.clone();
    let bar_cloned = bar.clone();
}

struct TestDeref {
    name: String,
    value: HashMap<i32, i32>,
}

impl TestDeref {
    fn test(&self) -> &str {
        // 返回self.name的引用
        &self.name
    }

    fn testGen<'a, T>(&self, a: &'a T) -> &'a T {
        a
    }
}

impl Index<i32> for TestDeref {
    type Output = i32;

    fn index(&self, index: i32) -> &Self::Output {
        self.value.get(&index).unwrap()
    }
}

trait Trait {}

fn foo2<X>(t: X)
where
    X: Trait,
{
}

impl Trait for &i32 {}
impl Trait for &mut i32 {}

fn get_memory_location() -> (usize, usize) {
    // “Hello World” 是字符串字面量，因此它的生命周期是 `'static`.
    // 但持有它的变量 `string` 的生命周期就不一样了，它完全取决于变量作用域，对于该例子来说，也就是当前的函数范围
    let string = "Hello World!";
    let pointer = string.as_ptr() as usize;
    let length = string.len();
    (pointer, length)
    // `string` 在这里被 drop 释放
    // 虽然变量被释放，无法再被访问，但是数据依然还会继续存活
}
fn get_str_at_location(pointer: usize, length: usize) -> &'static str {
    // 使用裸指针需要 `unsafe{}` 语句块
    unsafe { from_utf8_unchecked(from_raw_parts(pointer as *const u8, length)) }
}
fn read_username_from_file() -> result::Result<(), io::Error> {
    // read_to_string是定义在std::io中的方法，因此需要在上面进行引用
    let b = fs::read_to_string("hello.txt")?;
    Ok(())
}

// fn read() -> result::Result<(), Error> {
//     let b = File::open("path");
//     let er = Error {};
//     return Result::Err(er);
// }
// impl Add for String {
//     type Output = String;

//     fn add(self, rhs: Self) -> Self::Output {
//         todo!()
//     }
// }
// pub trait Draw {
//     fn draw(&self);
// }

pub trait Say: Display {
    fn bark(&self) {
        println!("实现我 bark");
    }
}

pub struct Cat {
    name: String,
}

impl Cat {
    fn bark(&self) {
        println!("{} wwww is barking", self.name);
    }
}
impl Display for Cat {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        write!(f, "cat")
    }
}

impl Say for Cat {
    fn bark(&self) {
        println!("{} is barking", self.name);
    }
}
