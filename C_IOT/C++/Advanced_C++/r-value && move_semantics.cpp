// r-value, move semantics

// - motivation:
// to avoid copy of a huge data

operator && => to have the constructor
technique to move pointer instead of data
std::move => static_cast<...> == cast of referense from lvalue to rvalue
