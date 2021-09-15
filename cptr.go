package numbers

/*
	int sum(int n,int *a){
		int sum=0;
		for (int i=0;i<n;i++) {
			sum+=*a;
			a+=sizeof(int)/2;
		}

		return sum;
	}

	int dot(int n, int *a, int *b){
		int mul =0;
		for(int i=0;i<n;i++){
			mul+=(*a)*(*b);
			a+=sizeof(int)/2;
			b+=sizeof(int)/2;
		}
		return mul;
	}

*/
import "C"

import (
	"unsafe"
)

func Sum(array []int) int {
	return int(C.sum(C.int(len(array)), (*C.int)(unsafe.Pointer(&array[0]))))
}

func Dot(a, b []int) int {
	return int(C.dot(C.int(len(a)), (*C.int)(unsafe.Pointer(&a[0])), (*C.int)(unsafe.Pointer(&b[0]))))
}
