resource "dblab_clone" "clone_1" {
  protected         = false
  id                = "clone_name"
  database_username = "myuser"
  database_password = "mypassword"
}
