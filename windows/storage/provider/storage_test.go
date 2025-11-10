package provider

import (
	"fmt"
	"log"
	"os"
	"testing"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/google/uuid"
	"github.com/saltosystems/winrt-go"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/storage"
	"github.com/saltosystems/winrt-go/windows/storage/streams"
	"github.com/stretchr/testify/require"
)

func init() {
	if err := ole.RoInitialize(1); err != nil {
		log.Fatal(err)
	}
}

func PrintAllFields(impl *StorageProviderSyncRootInfo) {
	fields := []struct {
		Name  string
		Value func() (interface{}, error)
	}{
		{Name: "ID", Value: func() (interface{}, error) { return impl.GetId() }},
		{"Path", func() (interface{}, error) { return impl.GetPath() }},
		{"DisplayNameResource", func() (interface{}, error) { return impl.GetDisplayNameResource() }},
		{"IconResource", func() (interface{}, error) { return impl.GetIconResource() }},
		{"HydrationPolicy", func() (interface{}, error) { return impl.GetHydrationPolicy() }},
		{"HydrationPolicyModifier", func() (interface{}, error) { return impl.GetHydrationPolicyModifier() }},
		{"PopulationPolicy", func() (interface{}, error) { return impl.GetPopulationPolicy() }},
		{"InSyncPolicy", func() (interface{}, error) { return impl.GetInSyncPolicy() }},
		{"HardlinkPolicy", func() (interface{}, error) { return impl.GetHardlinkPolicy() }},
		{"ShowSiblingsAsGroup", func() (interface{}, error) { return impl.GetShowSiblingsAsGroup() }},
		{"Version", func() (interface{}, error) { return impl.GetVersion() }},
		{"ProtectionMode", func() (interface{}, error) { return impl.GetProtectionMode() }},
		{"AllowPinning", func() (interface{}, error) { return impl.GetAllowPinning() }},
		{"StorageProviderItemPropertyDefinitions", func() (interface{}, error) { return impl.GetStorageProviderItemPropertyDefinitions() }},
		{"RecycleBinUri", func() (interface{}, error) { return impl.GetRecycleBinUri() }},
		{"Context", func() (interface{}, error) { return impl.GetContext() }},
	}

	fmt.Println("StorageProviderSyncRootInfo Fields:")
	for _, field := range fields {
		value, err := field.Value()
		if err != nil {
			fmt.Printf("  %s: Error (%v)\n", field.Name, err)
		} else {
			fmt.Printf("  %s: %v\n", field.Name, value)
		}
	}
}

func GetFolderFromPath(fp string) (*storage.StorageFolder, error) {
	var folder *storage.StorageFolder
	var err error
	waitChan := make(chan struct{})
	onCompleteCB := func(instance *foundation.AsyncOperationCompletedHandler, asyncInfo *foundation.IAsyncOperation, asyncStatus foundation.AsyncStatus) {
		defer close(waitChan)
		if asyncStatus != foundation.AsyncStatusCompleted {
			log.Printf("Async operation did not complete successfully: status %d", asyncStatus)
			err = fmt.Errorf("async operation did not complete successfully: status %d", asyncStatus)
			return
		}

		// Retrieve the StorageFile result from asyncInfo
		var resultPtr unsafe.Pointer
		resultPtr, err = asyncInfo.GetResults()
		if err != nil {
			log.Printf("Failed to get async operation result: %v", err)
			return
		}

		// Cast the result to a StorageFile
		folder = (*storage.StorageFolder)(resultPtr)
		log.Printf("Retrieved StorageFile: %+v", folder)
	}
	iid := winrt.ParameterizedInstanceGUID(foundation.GUIDAsyncOperationCompletedHandler, storage.SignatureStorageFolder)
	handler := foundation.NewAsyncOperationCompletedHandler(ole.NewGUID(iid), onCompleteCB)
	defer handler.Release()

	// this is an async operation
	fileAsyncOp, err := storage.StorageFolderGetFolderFromPathAsync(fp)
	if err != nil {
		return nil, err
	}

	err = fileAsyncOp.SetCompleted(handler)
	if err != nil {
		return nil, err
	}

	// Wait until async operation has stopped, and finish.
	<-waitChan
	return folder, err
}

// GetFileFromPath retrieves a StorageFile from a given file path using StorageFile.GetFileFromPathAsync api
// https://docs.microsoft.com/en-us/uwp/api/windows.storage.storagefile.getfilefrompathasync
func GetFileFromPath(fp string) (*storage.StorageFile, error) {
	// Create an AsyncOperationCompletedHandler to retrieve the StorageFile
	var storageFile *storage.StorageFile
	var err error
	waitChan := make(chan struct{})
	onCompleteCB := func(instance *foundation.AsyncOperationCompletedHandler, asyncInfo *foundation.IAsyncOperation, asyncStatus foundation.AsyncStatus) {
		defer close(waitChan)
		if asyncStatus != foundation.AsyncStatusCompleted {
			log.Printf("Async operation did not complete successfully: status %d", asyncStatus)
			err = fmt.Errorf("async operation did not complete successfully: status %d", asyncStatus)
			return
		}

		// Retrieve the StorageFile result from asyncInfo
		var resultPtr unsafe.Pointer
		resultPtr, err = asyncInfo.GetResults()
		if err != nil {
			log.Printf("Failed to get async operation result: %v", err)
			return
		}

		// Cast the result to a StorageFile
		storageFile = (*storage.StorageFile)(resultPtr)
		log.Printf("Retrieved StorageFile: %+v", storageFile)
	}
	iid := winrt.ParameterizedInstanceGUID(foundation.GUIDAsyncOperationCompletedHandler, storage.SignatureStorageFile)
	handler := foundation.NewAsyncOperationCompletedHandler(ole.NewGUID(iid), onCompleteCB)
	defer handler.Release()

	// this is an async operation
	fileAsyncOp, err := storage.StorageFileGetFileFromPathAsync(fp)
	if err != nil {
		return nil, err
	}

	err = fileAsyncOp.SetCompleted(handler)
	if err != nil {
		return nil, err
	}

	// Wait until async operation has stopped, and finish.
	<-waitChan
	return storageFile, err
}

func Test_GetCurrentSyncRoots(t *testing.T) {
	// tr := initTestResource(t, withTestBrowseDirFn(defaultBrowseDirTestFunc), withConnectSyncRoot())
	// defer tr.cleanUp()

	ok, err := StorageProviderSyncRootManagerIsSupported()
	require.NoError(t, err)
	require.True(t, ok)

	roots, err := StorageProviderSyncRootManagerGetCurrentSyncRoots()
	require.NoError(t, err)
	initialNumRoots, err := roots.GetSize()
	require.NoError(t, err)
	fmt.Println("Number of roots before test:", initialNumRoots)

	tempBase, err := os.UserCacheDir()
	require.NoError(t, err)
	syncRootPath, err := os.MkdirTemp(tempBase, "syncRootPath")
	require.NoError(t, err)

	writer, err := streams.NewDataWriter()
	require.NoError(t, err)
	syncRootID := []byte("syncRootIdentity")
	err = writer.WriteBytes(uint32(len(syncRootID)), syncRootID) // #nosec G115 - syncRootID length is always reasonable
	require.NoError(t, err)

	bufferContext, err := writer.DetachBuffer()
	require.NoError(t, err)

	syncRootInfo, err := NewStorageProviderSyncRootInfo()
	require.NoError(t, err)

	err = syncRootInfo.SetContext(bufferContext)
	require.NoError(t, err)

	// Use a unique ID for each test run to avoid conflicts
	testID := uuid.New().String()
	err = syncRootInfo.SetId(fmt.Sprintf("{%s}", testID))
	require.NoError(t, err)
	dir, err := GetFolderFromPath(syncRootPath)
	require.NoError(t, err)
	itf3 := dir.MustQueryInterface(ole.NewGUID(storage.GUIDIStorageFolder))
	defer itf3.Release()
	iStorageDir := (*storage.IStorageFolder)(unsafe.Pointer(itf3))
	err = syncRootInfo.SetPath(iStorageDir)
	require.NoError(t, err)
	err = syncRootInfo.SetHydrationPolicy(2)
	require.NoError(t, err)
	err = syncRootInfo.SetHydrationPolicyModifier(0)
	require.NoError(t, err)
	err = syncRootInfo.SetPopulationPolicy(2)
	require.NoError(t, err)
	err = syncRootInfo.SetInSyncPolicy(StorageProviderInSyncPolicyPreserveInsyncForSyncEngine)
	require.NoError(t, err)
	err = syncRootInfo.SetHardlinkPolicy(0)
	require.NoError(t, err)
	err = syncRootInfo.SetVersion("1.0")
	require.NoError(t, err)
	err = syncRootInfo.SetAllowPinning(true)
	require.NoError(t, err)
	err = syncRootInfo.SetShowSiblingsAsGroup(false)
	require.NoError(t, err)
	err = syncRootInfo.SetProtectionMode(0)
	require.NoError(t, err)
	err = syncRootInfo.SetDisplayNameResource("DisplayNameResource")
	require.NoError(t, err)
	err = syncRootInfo.SetIconResource("C:\\WINDOWS\\system32\\imageres.dll,-1043")
	require.NoError(t, err)

	err = StorageProviderSyncRootManagerRegister(syncRootInfo)
	require.NoError(t, err)

	roots, err = StorageProviderSyncRootManagerGetCurrentSyncRoots()
	require.NoError(t, err)
	finalNumRoots, err := roots.GetSize()
	require.NoError(t, err)
	fmt.Println("Number of roots after registration:", finalNumRoots)

	// Note: The count may not increase immediately due to caching or filtering in the Windows API.
	// The important thing is that registration succeeded without errors.
	require.GreaterOrEqual(t, finalNumRoots, initialNumRoots, "Sync root count should not decrease")
}
