const print = @import("std").debug.print;
const os = @import("std").os;

pub fn sort() !void {

    // check os
    print("{}\n", .{});

    // get downloads dir
    const downloadDir = os.chdir("~/Downloads");

    // filter file
}

pub fn main() !void {
    try sort();
}
