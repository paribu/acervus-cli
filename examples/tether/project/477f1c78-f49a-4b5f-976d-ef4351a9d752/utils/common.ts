import { BigInt } from "../lib/bigint";

// Load string from memory via given pointer.
// It reads bytes from memory and converts them into chars one by one.
export function loadString(ptr: i32): string {
	let tmp = ptr;
	let str = "";
	let char: u8;
	// Null terminator
	while ((char = load<u8>(ptr)) != 0) {
		str += String.fromCharCode(char);
		ptr++;
	}
	freeMemory(tmp);
	return str;
}

// Write given string into memory and return the pointer.
export function putString(str: string): i32 {
	let bytes = String.UTF8.encode(str);
	let ptr = allocate(bytes.byteLength + 1); // +1 for null terminator
	for (let i = 0; i < bytes.byteLength; i++) {
		store<u8>(ptr + i, load<u8>(changetype<usize>(bytes) + i));
	}
	store<u8>(ptr + bytes.byteLength, 0); // Null terminator
	return ptr;
}

export function allocate(size: i32): i32 {
	const ptr = <i32>__alloc(size);
 	if (ptr > i32.MAX_VALUE) {
        throw new Error("Pointer value exceeds i32 range");
    }
	return ptr;
}

export function convertToString<T>(value: T): string {
	if (value instanceof Array) {
		let result = "[";
		for (let i = 0; i < value.length; i++) {
				result += convertToString(value[i]);
			if (i < value.length - 1) {
				result += ",";
			}
		}
		result += "]";
		return result;
	}

	if (isBoolean(value)) {
        return value ? "true" : "false";
    }

	if (value instanceof BigInt) {
		return (value as BigInt).toString();
	}

	// @ts-ignore
	return value.toString();
}
