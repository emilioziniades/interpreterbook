let map = fn(arr, f) {
	let iter = fn(arr, acc) {
		if (len(arr) == 0) {
			acc
		} else {
			iter(rest(arr), push(acc, f(first(arr))))

		}
	};

	iter(arr, []);
};

let a = [1,2,3,4];
let double = fn(x) { x + x};
map(a, double)
