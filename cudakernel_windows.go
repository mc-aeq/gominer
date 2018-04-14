// Copyright (c) 2016 The Decred developers.
// Copyright (c) 2018 The Aequator developers.
// +build cuda

package main

import (
	"syscall"
	"unsafe"

	"github.com/barnex/cuda5/cu"
)

var (
	//kernelDll           = syscall.MustLoadDLL("aequator.dll")
	kernelDll               = syscall.MustLoadDLL("aequator.dll")
	precomputeTableProcAddr = kernelDll.MustFindProc("aequator_cpu_setBlock_52").Addr()
	kernelProcAddr          = kernelDll.MustFindProc("aequator_hash_nonce").Addr()
)

func cudaPrecomputeTable(input *[192]byte) {
	syscall.Syscall(precomputeTableProcAddr, 1, uintptr(unsafe.Pointer(input)), 0, 0)
}

func cudaInvokeKernel(gridx, blockx, threads uint32, startNonce uint32, nonceResults cu.DevicePtr, targetHigh uint32) {
	syscall.Syscall6(kernelProcAddr, 6, uintptr(gridx), uintptr(blockx), uintptr(threads),
		uintptr(startNonce), uintptr(nonceResults), uintptr(targetHigh))
}
