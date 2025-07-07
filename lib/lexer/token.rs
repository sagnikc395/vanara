use std::iter::Enumerate;
use std::ops::{Range, RangeFrom, RangeFull, RangeTo};

pub enum Token {
    Illegal,
    EOF,
    //identifier and literals
    Ident(String),
}
