// Databricks FileSystem (DBFS) API
//
// We strongly recommend using clients created via
// [github.com/databricks/databricks-sdk-go/workspaces.New] to simplify
// configuration experience.
//
// Please use the high-level [DbfsAPI.Open] and [DbfsAPI.Overwrite] methods
// to work with remote files through Go's [io] interfaces. The return value
// of [DbfsAPI.Open] implements the [io.Reader] and [io.WriterTo] interfaces.
// The [io.WriterTo] interface is used by [io.Copy] and maximizes throughput by
// reading data with the DBFS maximum read chunk size of 1MB.
//
// Internally, these methods wrap the low level [DbfsAPI.Create],
// [DbfsAPI.Close], [DbfsAPI.Read], and [DbfsAPI.AddBlock] methods:
//
//	upload, _ := os.Open("/path/to/local/file.ext")
//	_ = w.Dbfs.Overwrite(ctx, "/path/to/remote/file", upload)
//
//	download, _ := os.Create("/path/to/local")
//	remote, _ := w.Dbfs.Open(ctx, "/path/to/remote")
//	_ = io.Copy(download, remote)
//
// Moving files:
//
//	err := w.Dbfs.Move(ctx, dbfs.Move{
//		SourcePath:      "/remote/src/path",
//		DestinationPath: "/remote/dst/path",
//	})
//
// Creating directories:
//
//	w.Dbfs.MkdirsByPath(ctx, "/remote/dir/path")
package dbfs
