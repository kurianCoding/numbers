package main

/*
	#include<stdio.h>

	int sum(int n,int *a){
		int sum=0;
		for (int i=0;i<n;i++) {
			sum+=*a;
			a+=sizeof(int)/2;
		}

		return sum;
	}

*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Sum(int []array) {
	fmt.Println(C.sum(C.int(len(array)), (*C.int)(unsafe.Pointer(&array[0]))))
}
