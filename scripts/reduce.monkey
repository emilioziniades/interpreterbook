let reduce = fn(arr, initial, f) {
    let iter = fn(arr, acc) {
        if (len(arr) == 0 ) {
            acc
        } else {
            iter(rest(arr), f(acc, first(arr)))
        }
    }

    iter(arr, initial)
}

let sum = fn(arr) {
    reduce(arr, 0, fn(x, y) { x + y})
}

sum([1,2,3,4,5])
