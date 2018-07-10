package volume

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
	"fmt"

	"github.com/alibaba/pouch/storage/volume/driver"
	"github.com/alibaba/pouch/storage/volume/types"
)

func createVolumeCore(root string) (*Core, error) {
	cfg := Config{
		VolumeMetaPath: path.Join(root, "volume.db"),
	}

	return NewCore(cfg)
}

func TestCreateVolume(t *testing.T) {
	volumeDriverName := "fake1"

	dir, err := ioutil.TempDir("", "TestCreateVolume")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// create volume core
	core, err := createVolumeCore(dir)
	if err != nil {
		t.Fatal(err)
	}

	driver.Register(driver.NewFakeDriver(volumeDriverName))
	defer driver.Unregister(volumeDriverName)

	v, err := core.CreateVolume(types.VolumeID{Name: "test1", Driver: volumeDriverName})
	if err != nil {
		t.Fatalf("create volume error: %v", err)
	}

	if v.Name != "test1" {
		t.Fatalf("expect volume name is %s, but got %s", "test1", v.Name)
	}
	if v.Driver() != volumeDriverName {
		t.Fatalf("expect volume driver is %s, but got %s", volumeDriverName, v.Driver())
	}

	_, err = core.CreateVolume(types.VolumeID{Name: "none", Driver: "none"})
	if err == nil {
		t.Fatal("expect get driver not found error, but err is nil")
	}
}

func TestGetVolume(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestGetVolume")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// create volume core
	core, err := createVolumeCore(dir)
	if err != nil {
		t.Fatal(err)
	}

	volumeDriverName := "fake1"
	driver.Register(driver.NewFakeDriver(volumeDriverName))
	defer driver.Unregister(volumeDriverName)

	// Test1
	v, err := core.GetVolume(types.VolumeID{Name: "test1", Driver: volumeDriverName})
	// TODO nil
	fmt.Println(v)
	if err != nil {
		t.Fatalf("get volume error: %v", err)
	}

	if v.Name != "test1" {
		t.Fatalf("expect volume name is %s, but got %s", "test1", v.Name)
	}
	if v.Driver() != volumeDriverName {
		t.Fatalf("expect volume driver is %s, but got %s", volumeDriverName, v.Driver())
	}

	// Test2
	_, err = core.GetVolume(types.VolumeID{Name: "test2", Driver: "none"})
	if err == nil {
		t.Fatal("expect get driver not found error, but err is nil")
	}
}

func TestListVolumes(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestListVolumes")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// create volume core
	core, err := createVolumeCore(dir)
	if err != nil {
		t.Fatal(err)
	}

	volumeDriverName := "fake1"
	driver.Register(driver.NewFakeDriver(volumeDriverName))
	defer driver.Unregister(volumeDriverName)

	// Test1
	v, err := core.ListVolumes(nil)
	if err != nil {
		t.Fatal("expect get driver not found error, but err is nil")
	}
}

func TestListVolumeName(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestListVolumeName")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// create volume core
	core, err := createVolumeCore(dir)
	if err != nil {
		t.Fatal(err)
	}

	volumeDriverName := "fake1"
	driver.Register(driver.NewFakeDriver(volumeDriverName))
	defer driver.Unregister(volumeDriverName)

	// Test1
	v, err := core.ListVolumeName(nil)
	if err != nil {
		t.Fatal("expect get driver not found error, but err is nil")
	}
}

func TestRemoveVolume(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestRemoveVolume")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// create volume core
	core, err := createVolumeCore(dir)
	if err != nil {
		t.Fatal(err)
	}

	volumeDriverName := "fake1"
	driver.Register(driver.NewFakeDriver(volumeDriverName))
	defer driver.Unregister(volumeDriverName)

	// Test1
	err := core.RemoveVolume(types.VolumeID{Name: "test1", Driver: volumeDriverName})
	if err == nil {
		t.Fatal("expect get driver not found error, but err is nil")
	}
}

func TestVolumePath(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestVolumePath")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// create volume core
	core, err := createVolumeCore(dir)
	if err != nil {
		t.Fatal(err)
	}

	volumeDriverName := "fake1"
	driver.Register(driver.NewFakeDriver(volumeDriverName))
	defer driver.Unregister(volumeDriverName)

	// Test1
	s, err := core.VolumePath(types.VolumeID{Name: "test1", Driver: volumeDriverName})
	// TODO nil
	fmt.Println(s)
	if err != nil {
		t.Fatalf("VolumePath error: %v", err)
	}
}

func TestAttachVolume(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestAttachVolume")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// create volume core
	core, err := createVolumeCore(dir)
	if err != nil {
		t.Fatal(err)
	}

	volumeDriverName := "fake1"
	driver.Register(driver.NewFakeDriver(volumeDriverName))
	defer driver.Unregister(volumeDriverName)

	// Test1
	m := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	v, err := core.AttachVolume(types.VolumeID{Name: "test1", Driver: volumeDriverName}, m)
	// TODO nil
	fmt.Println(v)
	if err != nil {
		t.Fatalf("AttachVolume error: %v", err)
	}
}

func TestDetachVolume(t *testing.T) {
	// TODO
}
