use rand::{Rng, rng};

use std::cell::{Cell, RefCell};

use std::rc::Rc;
use std::result::Result;
use std::sync::Mutex;
use std::thread;
use std::{collections::HashMap, sync::Arc};

fn main() {
    let a: Rc<HashMap<i32, i32>> = Rc::new(HashMap::new()); //只读计数 
    let b: Arc<HashMap<i32, i32>> = Arc::new(HashMap::new()); //线程只读计数

    let d: Cell<HashMap<i32, i32>> = Cell::new(HashMap::new()); //copy 读写计数

    let d = Cell::new(String::from("2222")); //  读写计数
    // let e = d.get(); //没得Copy
    // d.set(String::from("value"));

    let e = RefCell::new(HashMap::new());
    e.borrow_mut().insert("k", "v");
    e.borrow_mut().insert("k2", "v2");

    println!("{e:?}");
}

fn test() {
    let mut handles = vec![];
    let num_threads = 3;
    let adds_per_thread = 3;

    let mut xx = HashMap::new();
    xx.insert(1, 2);

    let ht = Arc::new(xx);

    // b.insert(1, 2);
    // ht.get().insert("k", "2");

    for i in 0..num_threads {
        let ht = Arc::clone(&ht);

        let handle = thread::spawn(move || {
            for j in 0..adds_per_thread {
                let key = rng().random::<u32>();
                let value = rng().random::<u32>();
            }
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }
}
