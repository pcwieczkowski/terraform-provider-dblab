package provider

// func TestAccSnapshotsDataSource(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { testAccPreCheck(t) },
// 		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Read testing
// 			{
// 				Config: testAccSnapshotsDataSourceConfig,
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr("data.dblab_snapshots.test", "pool", "test"),
// 				),
// 			},
// 		},
// 	})
// }

// const testAccSnapshotsDataSourceConfig = providerConfig + `
// data "dblab_snapshots" "test" {
//   pool = "test"
// }
// `
