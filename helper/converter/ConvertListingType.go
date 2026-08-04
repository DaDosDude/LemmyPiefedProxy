package converter

import (
	"LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/model/piefed"
)

func ConvertListingType(in piefed.ListingType) lemmy.ListingType {
	if in == piefed.ListingTypePopular {
		return lemmy.ListingTypeAll
	}

	return lemmy.ListingType(in)
}

func ReverseConvertListingType(in lemmy.ListingType) piefed.ListingType {
	return piefed.ListingType(in)
}
