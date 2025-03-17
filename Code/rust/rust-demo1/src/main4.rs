struct Test<'a> {
    name: String,
    ptr_name: &'a str,
}

fn main() {
    let s = String::from("hello");
    let mut t = Test {
        name: s,
        ptr_name: "hello",
    };
    t.ptr_name = &t.ptr_name[1..];
    // println!("{}", t.ptr_name);
    // println!("{}", t.name);

    let mut tricky = WhatAboutThis {
        name: "Annabelle".to_string(),
        nickname: None,
    };
    tricky.tie_the_knot();
    // cannot borrow `tricky` as immutable because it is also borrowed as mutable
    // println!("{:?}", tricky);
}
#[derive(Debug)]
struct WhatAboutThis<'a> {
    name: String,
    nickname: Option<&'a str>,
}

impl<'a> WhatAboutThis<'a> {
    fn tie_the_knot(&'a mut self) {
        self.nickname = Some(&self.name[..4]);
    }
}
