use std::{
    sync::{Arc, Condvar, Mutex, Once, mpsc::channel},
    thread,
};

static INIT: Once = Once::new();
fn main() {
    let aa = Arc::new((Mutex::new(false), Condvar::new()));

    let bb = aa.clone();

    let handle = thread::spawn(move || {
        let (lock, cvar) = &*bb;
        let mut started = lock.lock().unwrap();
        *started = true;
        cvar.notify_one();
    });
    let (lock, cvar) = &*aa;
    let mut started = lock.lock().unwrap();
    while !*started {
        let b = cvar.wait(started).unwrap(); // 等待通知
        started = b;
    }

    let a = thread::spawn(|| {
        println!("a,   start   ");
        INIT.call_once(|| println!("我只会被调用一次任何时候 a, world!"));
        println!("a,   end   ");
    });
    let b = thread::spawn(|| {
        println!("b,    start  ");
        INIT.call_once(|| println!("我只会被调用一次任何时候 b, world!"));
        println!("b,   end   ");
    });
    let c = thread::spawn(|| {
        println!("c,   start   ");
        INIT.call_once(|| println!("我只会被调用一次任何时候b , world!"));
        println!("c,   end   ");
    });
    a.join().unwrap();
    b.join().unwrap();
    c.join().unwrap();
    println!("完毕");

    let (ts, re) = channel();

    thread::spawn(move || {
        for x in 1..100 {
            ts.send(x).unwrap();
            println!("send {x}");
        }
    });
    let g = loop {
        let x = re.recv(); // 阻塞
        let b = match x {
            Ok(c) => c,
            Err(x) => break 22222,
        };
        println!("recv {b}");
    };

    println!("要输出22222 {g:?}");
    assert_eq!(g, 22222);

    let (tx, rx) = channel();
    thread::spawn(move || {
        tx.send("hello").unwrap();
    });
    thread::sleep(std::time::Duration::from_secs(1));
    let s = rx.try_recv(); //不阻塞
    println!("要输出hello {s:?}");

    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let s = String::from("我，飞走咯!");
        tx.send(s).unwrap();
        println!("val is {}", s);
    });

    let received = rx.recv().unwrap();
    println!("Got: {}", received);
}
