create table if not exists users (
  id integer primary key,
  screen_name text,
  display_name text,
  password text,
  email text,
  created_at datetime,
  updated_at datetime
)
