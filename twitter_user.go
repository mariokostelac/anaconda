package twitter

type TwitterUser struct {
	Statuses_count                     *float64
	Contributors_enabled               *bool
	Friends_count                      *float64
	Geo_enabled                        *bool
	Description                        *string
	Profile_sidebar_border_color       *string
	Screen_name                        *string
	Listed_count                       *float64
	Followers_count                    *float64
	Location                           *string
	Profile_background_image_url       *string
	Name                               *string
	Default_profile_image              *bool
	Profile_image_url_https            *string
	Notifications                      *bool
	Protected                          *bool
	Id_str                             *string
	Profile_background_color           *string
	Created_at                         *string
	Default_profile                    *bool
	Url                                *string
	Id                                 *float64
	Verified                           *bool
	Profile_link_color                 *string
	Profile_image_url                  *string
	Profile_use_background_image       *bool
	Favourites_count                   *float64
	Profile_background_image_url_https *string
	Profile_sidebar_fill_color         *string
	Is_translator                      *bool
	Follow_request_sent                *bool
	Following                          *bool
	Profile_background_tile            *bool
	Show_all_inline_media              *bool
	Profile_text_color                 *string
	Lang                               *string
	Entities                           *struct {
		Url struct {
			Urls []interface{}
		}
		Description struct {
			Urls []interface{}
		}
	}
}
