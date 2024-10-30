// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package bluetooth

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const SignatureBluetoothLEDevice string = "rc(Windows.Devices.Bluetooth.BluetoothLEDevice;{b5ee2f7b-4ad8-4642-ac48-80a0b500e887})"

type BluetoothLEDevice struct {
	ole.IUnknown
}

func (impl *BluetoothLEDevice) GetConnectionStatus() (BluetoothConnectionStatus, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEDevice))
	defer itf.Release()
	v := (*iBluetoothLEDevice)(unsafe.Pointer(itf))
	return v.GetConnectionStatus()
}

func (impl *BluetoothLEDevice) AddConnectionStatusChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEDevice))
	defer itf.Release()
	v := (*iBluetoothLEDevice)(unsafe.Pointer(itf))
	return v.AddConnectionStatusChanged(handler)
}

func (impl *BluetoothLEDevice) RemoveConnectionStatusChanged(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEDevice))
	defer itf.Release()
	v := (*iBluetoothLEDevice)(unsafe.Pointer(itf))
	return v.RemoveConnectionStatusChanged(token)
}

func (impl *BluetoothLEDevice) GetGattServicesAsync() (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEDevice3))
	defer itf.Release()
	v := (*iBluetoothLEDevice3)(unsafe.Pointer(itf))
	return v.GetGattServicesAsync()
}

func (impl *BluetoothLEDevice) GetGattServicesWithCacheModeAsync(cacheMode BluetoothCacheMode) (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEDevice3))
	defer itf.Release()
	v := (*iBluetoothLEDevice3)(unsafe.Pointer(itf))
	return v.GetGattServicesWithCacheModeAsync(cacheMode)
}

func (impl *BluetoothLEDevice) GetBluetoothDeviceId() (*BluetoothDeviceId, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEDevice4))
	defer itf.Release()
	v := (*iBluetoothLEDevice4)(unsafe.Pointer(itf))
	return v.GetBluetoothDeviceId()
}

func (impl *BluetoothLEDevice) Close() error {
	itf := impl.MustQueryInterface(ole.NewGUID(foundation.GUIDIClosable))
	defer itf.Release()
	v := (*foundation.IClosable)(unsafe.Pointer(itf))
	return v.Close()
}

const GUIDiBluetoothLEDevice string = "b5ee2f7b-4ad8-4642-ac48-80a0b500e887"
const SignatureiBluetoothLEDevice string = "{b5ee2f7b-4ad8-4642-ac48-80a0b500e887}"

type iBluetoothLEDevice struct {
	ole.IInspectable
}

type iBluetoothLEDeviceVtbl struct {
	ole.IInspectableVtbl

	GetDeviceId                   uintptr
	GetName                       uintptr
	GetGattServices               uintptr
	GetConnectionStatus           uintptr
	GetBluetoothAddress           uintptr
	GetGattService                uintptr
	AddNameChanged                uintptr
	RemoveNameChanged             uintptr
	AddGattServicesChanged        uintptr
	RemoveGattServicesChanged     uintptr
	AddConnectionStatusChanged    uintptr
	RemoveConnectionStatusChanged uintptr
}

func (v *iBluetoothLEDevice) VTable() *iBluetoothLEDeviceVtbl {
	return (*iBluetoothLEDeviceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iBluetoothLEDevice) GetConnectionStatus() (BluetoothConnectionStatus, error) {
	var out BluetoothConnectionStatus
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetConnectionStatus,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out BluetoothConnectionStatus
	)

	if hr != 0 {
		return BluetoothConnectionStatusDisconnected, ole.NewError(hr)
	}

	return out, nil
}

func (v *iBluetoothLEDevice) AddConnectionStatusChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddConnectionStatusChanged,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iBluetoothLEDevice) RemoveConnectionStatusChanged(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveConnectionStatusChanged,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

const GUIDiBluetoothLEDevice2 string = "26f062b3-7aee-4d31-baba-b1b9775f5916"
const SignatureiBluetoothLEDevice2 string = "{26f062b3-7aee-4d31-baba-b1b9775f5916}"

type iBluetoothLEDevice2 struct {
	ole.IInspectable
}

type iBluetoothLEDevice2Vtbl struct {
	ole.IInspectableVtbl

	GetDeviceInformation    uintptr
	GetAppearance           uintptr
	GetBluetoothAddressType uintptr
}

func (v *iBluetoothLEDevice2) VTable() *iBluetoothLEDevice2Vtbl {
	return (*iBluetoothLEDevice2Vtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiBluetoothLEDevice3 string = "aee9e493-44ac-40dc-af33-b2c13c01ca46"
const SignatureiBluetoothLEDevice3 string = "{aee9e493-44ac-40dc-af33-b2c13c01ca46}"

type iBluetoothLEDevice3 struct {
	ole.IInspectable
}

type iBluetoothLEDevice3Vtbl struct {
	ole.IInspectableVtbl

	GetDeviceAccessInformation               uintptr
	RequestAccessAsync                       uintptr
	GetGattServicesAsync                     uintptr
	GetGattServicesWithCacheModeAsync        uintptr
	GetGattServicesForUuidAsync              uintptr
	GetGattServicesForUuidWithCacheModeAsync uintptr
}

func (v *iBluetoothLEDevice3) VTable() *iBluetoothLEDevice3Vtbl {
	return (*iBluetoothLEDevice3Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iBluetoothLEDevice3) GetGattServicesAsync() (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetGattServicesAsync,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iBluetoothLEDevice3) GetGattServicesWithCacheModeAsync(cacheMode BluetoothCacheMode) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetGattServicesWithCacheModeAsync,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(cacheMode),            // in BluetoothCacheMode
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiBluetoothLEDevice4 string = "2b605031-2248-4b2f-acf0-7cee36fc5870"
const SignatureiBluetoothLEDevice4 string = "{2b605031-2248-4b2f-acf0-7cee36fc5870}"

type iBluetoothLEDevice4 struct {
	ole.IInspectable
}

type iBluetoothLEDevice4Vtbl struct {
	ole.IInspectableVtbl

	GetBluetoothDeviceId uintptr
}

func (v *iBluetoothLEDevice4) VTable() *iBluetoothLEDevice4Vtbl {
	return (*iBluetoothLEDevice4Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iBluetoothLEDevice4) GetBluetoothDeviceId() (*BluetoothDeviceId, error) {
	var out *BluetoothDeviceId
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetBluetoothDeviceId,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out BluetoothDeviceId
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiBluetoothLEDevice5 string = "9d6a1260-5287-458e-95ba-17c8b7bb326e"
const SignatureiBluetoothLEDevice5 string = "{9d6a1260-5287-458e-95ba-17c8b7bb326e}"

type iBluetoothLEDevice5 struct {
	ole.IInspectable
}

type iBluetoothLEDevice5Vtbl struct {
	ole.IInspectableVtbl

	GetWasSecureConnectionUsedForPairing uintptr
}

func (v *iBluetoothLEDevice5) VTable() *iBluetoothLEDevice5Vtbl {
	return (*iBluetoothLEDevice5Vtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiBluetoothLEDevice6 string = "ca7190ef-0cae-573c-a1ca-e1fc5bfc39e2"
const SignatureiBluetoothLEDevice6 string = "{ca7190ef-0cae-573c-a1ca-e1fc5bfc39e2}"

type iBluetoothLEDevice6 struct {
	ole.IInspectable
}

type iBluetoothLEDevice6Vtbl struct {
	ole.IInspectableVtbl

	GetConnectionParameters              uintptr
	GetConnectionPhy                     uintptr
	RequestPreferredConnectionParameters uintptr
	AddConnectionParametersChanged       uintptr
	RemoveConnectionParametersChanged    uintptr
	AddConnectionPhyChanged              uintptr
	RemoveConnectionPhyChanged           uintptr
}

func (v *iBluetoothLEDevice6) VTable() *iBluetoothLEDevice6Vtbl {
	return (*iBluetoothLEDevice6Vtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiBluetoothLEDeviceStatics2 string = "5f12c06b-3bac-43e8-ad16-563271bd41c2"
const SignatureiBluetoothLEDeviceStatics2 string = "{5f12c06b-3bac-43e8-ad16-563271bd41c2}"

type iBluetoothLEDeviceStatics2 struct {
	ole.IInspectable
}

type iBluetoothLEDeviceStatics2Vtbl struct {
	ole.IInspectableVtbl

	BluetoothLEDeviceGetDeviceSelectorFromPairingState                             uintptr
	BluetoothLEDeviceGetDeviceSelectorFromConnectionStatus                         uintptr
	BluetoothLEDeviceGetDeviceSelectorFromDeviceName                               uintptr
	BluetoothLEDeviceGetDeviceSelectorFromBluetoothAddress                         uintptr
	BluetoothLEDeviceGetDeviceSelectorFromBluetoothAddressWithBluetoothAddressType uintptr
	BluetoothLEDeviceGetDeviceSelectorFromAppearance                               uintptr
	BluetoothLEDeviceFromBluetoothAddressWithBluetoothAddressTypeAsync             uintptr
}

func (v *iBluetoothLEDeviceStatics2) VTable() *iBluetoothLEDeviceStatics2Vtbl {
	return (*iBluetoothLEDeviceStatics2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func BluetoothLEDeviceFromBluetoothAddressWithBluetoothAddressTypeAsync(bluetoothAddress uint64, bluetoothAddressType BluetoothAddressType) (*foundation.IAsyncOperation, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Devices.Bluetooth.BluetoothLEDevice", ole.NewGUID(GUIDiBluetoothLEDeviceStatics2))
	if err != nil {
		return nil, err
	}
	v := (*iBluetoothLEDeviceStatics2)(unsafe.Pointer(inspectable))

	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().BluetoothLEDeviceFromBluetoothAddressWithBluetoothAddressTypeAsync,
		0,                             // this is a static func, so there's no this
		uintptr(bluetoothAddress),     // in uint64
		uintptr(bluetoothAddressType), // in BluetoothAddressType
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiBluetoothLEDeviceStatics string = "c8cf1a19-f0b6-4bf0-8689-41303de2d9f4"
const SignatureiBluetoothLEDeviceStatics string = "{c8cf1a19-f0b6-4bf0-8689-41303de2d9f4}"

type iBluetoothLEDeviceStatics struct {
	ole.IInspectable
}

type iBluetoothLEDeviceStaticsVtbl struct {
	ole.IInspectableVtbl

	BluetoothLEDeviceFromIdAsync               uintptr
	BluetoothLEDeviceFromBluetoothAddressAsync uintptr
	BluetoothLEDeviceGetDeviceSelector         uintptr
}

func (v *iBluetoothLEDeviceStatics) VTable() *iBluetoothLEDeviceStaticsVtbl {
	return (*iBluetoothLEDeviceStaticsVtbl)(unsafe.Pointer(v.RawVTable))
}

func BluetoothLEDeviceFromBluetoothAddressAsync(bluetoothAddress uint64) (*foundation.IAsyncOperation, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Devices.Bluetooth.BluetoothLEDevice", ole.NewGUID(GUIDiBluetoothLEDeviceStatics))
	if err != nil {
		return nil, err
	}
	v := (*iBluetoothLEDeviceStatics)(unsafe.Pointer(inspectable))

	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().BluetoothLEDeviceFromBluetoothAddressAsync,
		0,                             // this is a static func, so there's no this
		uintptr(bluetoothAddress),     // in uint64
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
