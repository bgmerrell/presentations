all_uuid_list = [ x.id for x in all_posts_uuid_list if ( x.type == "post" and not x.deleted ) ]
