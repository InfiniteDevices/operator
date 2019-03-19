// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/segments.proto

package common // import "google.golang.org/genproto/googleapis/ads/googleads/v1/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Segment only fields.
type Segments struct {
	// Ad network type.
	AdNetworkType enums.AdNetworkTypeEnum_AdNetworkType `protobuf:"varint,3,opt,name=ad_network_type,json=adNetworkType,proto3,enum=google.ads.googleads.v1.enums.AdNetworkTypeEnum_AdNetworkType" json:"ad_network_type,omitempty"`
	// Click type.
	ClickType enums.ClickTypeEnum_ClickType `protobuf:"varint,26,opt,name=click_type,json=clickType,proto3,enum=google.ads.googleads.v1.enums.ClickTypeEnum_ClickType" json:"click_type,omitempty"`
	// Resource name of the conversion action.
	ConversionAction *wrappers.StringValue `protobuf:"bytes,52,opt,name=conversion_action,json=conversionAction,proto3" json:"conversion_action,omitempty"`
	// Conversion action category.
	ConversionActionCategory enums.ConversionActionCategoryEnum_ConversionActionCategory `protobuf:"varint,53,opt,name=conversion_action_category,json=conversionActionCategory,proto3,enum=google.ads.googleads.v1.enums.ConversionActionCategoryEnum_ConversionActionCategory" json:"conversion_action_category,omitempty"`
	// Conversion action name.
	ConversionActionName *wrappers.StringValue `protobuf:"bytes,54,opt,name=conversion_action_name,json=conversionActionName,proto3" json:"conversion_action_name,omitempty"`
	// This segments your conversion columns by the original conversion and
	// conversion value vs. the delta if conversions were adjusted. False row has
	// the data as originally stated; While true row has the delta between data
	// now and the data as originally stated. Summing the two together results
	// post-adjustment data.
	ConversionAdjustment *wrappers.BoolValue `protobuf:"bytes,27,opt,name=conversion_adjustment,json=conversionAdjustment,proto3" json:"conversion_adjustment,omitempty"`
	// Conversion attribution event type.
	ConversionAttributionEventType enums.ConversionAttributionEventTypeEnum_ConversionAttributionEventType `protobuf:"varint,2,opt,name=conversion_attribution_event_type,json=conversionAttributionEventType,proto3,enum=google.ads.googleads.v1.enums.ConversionAttributionEventTypeEnum_ConversionAttributionEventType" json:"conversion_attribution_event_type,omitempty"`
	// An enum value representing the number of days between the impression and
	// the conversion.
	ConversionLagBucket enums.ConversionLagBucketEnum_ConversionLagBucket `protobuf:"varint,50,opt,name=conversion_lag_bucket,json=conversionLagBucket,proto3,enum=google.ads.googleads.v1.enums.ConversionLagBucketEnum_ConversionLagBucket" json:"conversion_lag_bucket,omitempty"`
	// An enum value representing the number of days between the impression and
	// the conversion or between the impression and adjustments to the conversion.
	ConversionOrAdjustmentLagBucket enums.ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket `protobuf:"varint,51,opt,name=conversion_or_adjustment_lag_bucket,json=conversionOrAdjustmentLagBucket,proto3,enum=google.ads.googleads.v1.enums.ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket" json:"conversion_or_adjustment_lag_bucket,omitempty"`
	// Date to which metrics apply.
	// yyyy-MM-dd format, e.g., 2018-04-17.
	Date *wrappers.StringValue `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	// Day of the week, e.g., MONDAY.
	DayOfWeek enums.DayOfWeekEnum_DayOfWeek `protobuf:"varint,5,opt,name=day_of_week,json=dayOfWeek,proto3,enum=google.ads.googleads.v1.enums.DayOfWeekEnum_DayOfWeek" json:"day_of_week,omitempty"`
	// Device to which metrics apply.
	Device enums.DeviceEnum_Device `protobuf:"varint,1,opt,name=device,proto3,enum=google.ads.googleads.v1.enums.DeviceEnum_Device" json:"device,omitempty"`
	// External conversion source.
	ExternalConversionSource enums.ExternalConversionSourceEnum_ExternalConversionSource `protobuf:"varint,55,opt,name=external_conversion_source,json=externalConversionSource,proto3,enum=google.ads.googleads.v1.enums.ExternalConversionSourceEnum_ExternalConversionSource" json:"external_conversion_source,omitempty"`
	// Resource name of the geo target constant that represents a city.
	GeoTargetCity *wrappers.StringValue `protobuf:"bytes,62,opt,name=geo_target_city,json=geoTargetCity,proto3" json:"geo_target_city,omitempty"`
	// Resource name of the geo target constant that represents a metro.
	GeoTargetMetro *wrappers.StringValue `protobuf:"bytes,63,opt,name=geo_target_metro,json=geoTargetMetro,proto3" json:"geo_target_metro,omitempty"`
	// Resource name of the geo target constant that represents a region.
	GeoTargetRegion *wrappers.StringValue `protobuf:"bytes,64,opt,name=geo_target_region,json=geoTargetRegion,proto3" json:"geo_target_region,omitempty"`
	// Hotel booking window in days.
	HotelBookingWindowDays *wrappers.Int64Value `protobuf:"bytes,6,opt,name=hotel_booking_window_days,json=hotelBookingWindowDays,proto3" json:"hotel_booking_window_days,omitempty"`
	// Hotel center ID.
	HotelCenterId *wrappers.Int64Value `protobuf:"bytes,7,opt,name=hotel_center_id,json=hotelCenterId,proto3" json:"hotel_center_id,omitempty"`
	// Hotel check-in date. Formatted as yyyy-MM-dd.
	HotelCheckInDate *wrappers.StringValue `protobuf:"bytes,8,opt,name=hotel_check_in_date,json=hotelCheckInDate,proto3" json:"hotel_check_in_date,omitempty"`
	// Hotel check-in day of week.
	HotelCheckInDayOfWeek enums.DayOfWeekEnum_DayOfWeek `protobuf:"varint,9,opt,name=hotel_check_in_day_of_week,json=hotelCheckInDayOfWeek,proto3,enum=google.ads.googleads.v1.enums.DayOfWeekEnum_DayOfWeek" json:"hotel_check_in_day_of_week,omitempty"`
	// Hotel city.
	HotelCity *wrappers.StringValue `protobuf:"bytes,10,opt,name=hotel_city,json=hotelCity,proto3" json:"hotel_city,omitempty"`
	// Hotel class.
	HotelClass *wrappers.Int32Value `protobuf:"bytes,11,opt,name=hotel_class,json=hotelClass,proto3" json:"hotel_class,omitempty"`
	// Hotel country.
	HotelCountry *wrappers.StringValue `protobuf:"bytes,12,opt,name=hotel_country,json=hotelCountry,proto3" json:"hotel_country,omitempty"`
	// Hotel date selection type.
	HotelDateSelectionType enums.HotelDateSelectionTypeEnum_HotelDateSelectionType `protobuf:"varint,13,opt,name=hotel_date_selection_type,json=hotelDateSelectionType,proto3,enum=google.ads.googleads.v1.enums.HotelDateSelectionTypeEnum_HotelDateSelectionType" json:"hotel_date_selection_type,omitempty"`
	// Hotel length of stay.
	HotelLengthOfStay *wrappers.Int32Value `protobuf:"bytes,14,opt,name=hotel_length_of_stay,json=hotelLengthOfStay,proto3" json:"hotel_length_of_stay,omitempty"`
	// Hotel state.
	HotelState *wrappers.StringValue `protobuf:"bytes,15,opt,name=hotel_state,json=hotelState,proto3" json:"hotel_state,omitempty"`
	// Hour of day as a number between 0 and 23, inclusive.
	Hour *wrappers.Int32Value `protobuf:"bytes,16,opt,name=hour,proto3" json:"hour,omitempty"`
	// Only used with feed item metrics.
	// Indicates whether the interaction metrics occurred on the feed item itself
	// or a different extension or ad unit.
	InteractionOnThisExtension *wrappers.BoolValue `protobuf:"bytes,49,opt,name=interaction_on_this_extension,json=interactionOnThisExtension,proto3" json:"interaction_on_this_extension,omitempty"`
	// Keyword criterion.
	Keyword *Keyword `protobuf:"bytes,61,opt,name=keyword,proto3" json:"keyword,omitempty"`
	// Month as represented by the date of the first day of a month. Formatted as
	// yyyy-MM-dd.
	Month *wrappers.StringValue `protobuf:"bytes,17,opt,name=month,proto3" json:"month,omitempty"`
	// Month of the year, e.g., January.
	MonthOfYear enums.MonthOfYearEnum_MonthOfYear `protobuf:"varint,18,opt,name=month_of_year,json=monthOfYear,proto3,enum=google.ads.googleads.v1.enums.MonthOfYearEnum_MonthOfYear" json:"month_of_year,omitempty"`
	// Partner hotel ID.
	PartnerHotelId *wrappers.StringValue `protobuf:"bytes,19,opt,name=partner_hotel_id,json=partnerHotelId,proto3" json:"partner_hotel_id,omitempty"`
	// Placeholder type. This is only used with feed item metrics.
	PlaceholderType enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,20,opt,name=placeholder_type,json=placeholderType,proto3,enum=google.ads.googleads.v1.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_type,omitempty"`
	// Aggregator ID of the product.
	ProductAggregatorId *wrappers.UInt64Value `protobuf:"bytes,28,opt,name=product_aggregator_id,json=productAggregatorId,proto3" json:"product_aggregator_id,omitempty"`
	// Bidding category (level 1) of the product.
	ProductBiddingCategoryLevel1 *wrappers.StringValue `protobuf:"bytes,56,opt,name=product_bidding_category_level1,json=productBiddingCategoryLevel1,proto3" json:"product_bidding_category_level1,omitempty"`
	// Bidding category (level 2) of the product.
	ProductBiddingCategoryLevel2 *wrappers.StringValue `protobuf:"bytes,57,opt,name=product_bidding_category_level2,json=productBiddingCategoryLevel2,proto3" json:"product_bidding_category_level2,omitempty"`
	// Bidding category (level 3) of the product.
	ProductBiddingCategoryLevel3 *wrappers.StringValue `protobuf:"bytes,58,opt,name=product_bidding_category_level3,json=productBiddingCategoryLevel3,proto3" json:"product_bidding_category_level3,omitempty"`
	// Bidding category (level 4) of the product.
	ProductBiddingCategoryLevel4 *wrappers.StringValue `protobuf:"bytes,59,opt,name=product_bidding_category_level4,json=productBiddingCategoryLevel4,proto3" json:"product_bidding_category_level4,omitempty"`
	// Bidding category (level 5) of the product.
	ProductBiddingCategoryLevel5 *wrappers.StringValue `protobuf:"bytes,60,opt,name=product_bidding_category_level5,json=productBiddingCategoryLevel5,proto3" json:"product_bidding_category_level5,omitempty"`
	// Brand of the product.
	ProductBrand *wrappers.StringValue `protobuf:"bytes,29,opt,name=product_brand,json=productBrand,proto3" json:"product_brand,omitempty"`
	// Channel of the product.
	ProductChannel enums.ProductChannelEnum_ProductChannel `protobuf:"varint,30,opt,name=product_channel,json=productChannel,proto3,enum=google.ads.googleads.v1.enums.ProductChannelEnum_ProductChannel" json:"product_channel,omitempty"`
	// Channel exclusivity of the product.
	ProductChannelExclusivity enums.ProductChannelExclusivityEnum_ProductChannelExclusivity `protobuf:"varint,31,opt,name=product_channel_exclusivity,json=productChannelExclusivity,proto3,enum=google.ads.googleads.v1.enums.ProductChannelExclusivityEnum_ProductChannelExclusivity" json:"product_channel_exclusivity,omitempty"`
	// Condition of the product.
	ProductCondition enums.ProductConditionEnum_ProductCondition `protobuf:"varint,32,opt,name=product_condition,json=productCondition,proto3,enum=google.ads.googleads.v1.enums.ProductConditionEnum_ProductCondition" json:"product_condition,omitempty"`
	// Resource name of the geo target constant for the country of the product.
	ProductCountry *wrappers.StringValue `protobuf:"bytes,33,opt,name=product_country,json=productCountry,proto3" json:"product_country,omitempty"`
	// Custom attribute 0 of the product.
	ProductCustomAttribute0 *wrappers.StringValue `protobuf:"bytes,34,opt,name=product_custom_attribute0,json=productCustomAttribute0,proto3" json:"product_custom_attribute0,omitempty"`
	// Custom attribute 1 of the product.
	ProductCustomAttribute1 *wrappers.StringValue `protobuf:"bytes,35,opt,name=product_custom_attribute1,json=productCustomAttribute1,proto3" json:"product_custom_attribute1,omitempty"`
	// Custom attribute 2 of the product.
	ProductCustomAttribute2 *wrappers.StringValue `protobuf:"bytes,36,opt,name=product_custom_attribute2,json=productCustomAttribute2,proto3" json:"product_custom_attribute2,omitempty"`
	// Custom attribute 3 of the product.
	ProductCustomAttribute3 *wrappers.StringValue `protobuf:"bytes,37,opt,name=product_custom_attribute3,json=productCustomAttribute3,proto3" json:"product_custom_attribute3,omitempty"`
	// Custom attribute 4 of the product.
	ProductCustomAttribute4 *wrappers.StringValue `protobuf:"bytes,38,opt,name=product_custom_attribute4,json=productCustomAttribute4,proto3" json:"product_custom_attribute4,omitempty"`
	// Item ID of the product.
	ProductItemId *wrappers.StringValue `protobuf:"bytes,39,opt,name=product_item_id,json=productItemId,proto3" json:"product_item_id,omitempty"`
	// Resource name of the language constant for the language of the product.
	ProductLanguage *wrappers.StringValue `protobuf:"bytes,40,opt,name=product_language,json=productLanguage,proto3" json:"product_language,omitempty"`
	// Merchant ID of the product.
	ProductMerchantId *wrappers.UInt64Value `protobuf:"bytes,41,opt,name=product_merchant_id,json=productMerchantId,proto3" json:"product_merchant_id,omitempty"`
	// Store ID of the product.
	ProductStoreId *wrappers.StringValue `protobuf:"bytes,42,opt,name=product_store_id,json=productStoreId,proto3" json:"product_store_id,omitempty"`
	// Title of the product.
	ProductTitle *wrappers.StringValue `protobuf:"bytes,43,opt,name=product_title,json=productTitle,proto3" json:"product_title,omitempty"`
	// Type (level 1) of the product.
	ProductTypeL1 *wrappers.StringValue `protobuf:"bytes,44,opt,name=product_type_l1,json=productTypeL1,proto3" json:"product_type_l1,omitempty"`
	// Type (level 2) of the product.
	ProductTypeL2 *wrappers.StringValue `protobuf:"bytes,45,opt,name=product_type_l2,json=productTypeL2,proto3" json:"product_type_l2,omitempty"`
	// Type (level 3) of the product.
	ProductTypeL3 *wrappers.StringValue `protobuf:"bytes,46,opt,name=product_type_l3,json=productTypeL3,proto3" json:"product_type_l3,omitempty"`
	// Type (level 4) of the product.
	ProductTypeL4 *wrappers.StringValue `protobuf:"bytes,47,opt,name=product_type_l4,json=productTypeL4,proto3" json:"product_type_l4,omitempty"`
	// Type (level 5) of the product.
	ProductTypeL5 *wrappers.StringValue `protobuf:"bytes,48,opt,name=product_type_l5,json=productTypeL5,proto3" json:"product_type_l5,omitempty"`
	// Quarter as represented by the date of the first day of a quarter.
	// Uses the calendar year for quarters, e.g., the second quarter of 2018
	// starts on 2018-04-01. Formatted as yyyy-MM-dd.
	Quarter *wrappers.StringValue `protobuf:"bytes,21,opt,name=quarter,proto3" json:"quarter,omitempty"`
	// Match type of the keyword that triggered the ad, including variants.
	SearchTermMatchType enums.SearchTermMatchTypeEnum_SearchTermMatchType `protobuf:"varint,22,opt,name=search_term_match_type,json=searchTermMatchType,proto3,enum=google.ads.googleads.v1.enums.SearchTermMatchTypeEnum_SearchTermMatchType" json:"search_term_match_type,omitempty"`
	// Position of the ad.
	Slot enums.SlotEnum_Slot `protobuf:"varint,23,opt,name=slot,proto3,enum=google.ads.googleads.v1.enums.SlotEnum_Slot" json:"slot,omitempty"`
	// Week as defined as Monday through Sunday, and represented by the date of
	// Monday. Formatted as yyyy-MM-dd.
	Week *wrappers.StringValue `protobuf:"bytes,24,opt,name=week,proto3" json:"week,omitempty"`
	// Year, formatted as yyyy.
	Year                 *wrappers.Int32Value `protobuf:"bytes,25,opt,name=year,proto3" json:"year,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Segments) Reset()         { *m = Segments{} }
func (m *Segments) String() string { return proto.CompactTextString(m) }
func (*Segments) ProtoMessage()    {}
func (*Segments) Descriptor() ([]byte, []int) {
	return fileDescriptor_segments_74ce948b3f70e9be, []int{0}
}
func (m *Segments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Segments.Unmarshal(m, b)
}
func (m *Segments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Segments.Marshal(b, m, deterministic)
}
func (dst *Segments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Segments.Merge(dst, src)
}
func (m *Segments) XXX_Size() int {
	return xxx_messageInfo_Segments.Size(m)
}
func (m *Segments) XXX_DiscardUnknown() {
	xxx_messageInfo_Segments.DiscardUnknown(m)
}

var xxx_messageInfo_Segments proto.InternalMessageInfo

func (m *Segments) GetAdNetworkType() enums.AdNetworkTypeEnum_AdNetworkType {
	if m != nil {
		return m.AdNetworkType
	}
	return enums.AdNetworkTypeEnum_UNSPECIFIED
}

func (m *Segments) GetClickType() enums.ClickTypeEnum_ClickType {
	if m != nil {
		return m.ClickType
	}
	return enums.ClickTypeEnum_UNSPECIFIED
}

func (m *Segments) GetConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.ConversionAction
	}
	return nil
}

func (m *Segments) GetConversionActionCategory() enums.ConversionActionCategoryEnum_ConversionActionCategory {
	if m != nil {
		return m.ConversionActionCategory
	}
	return enums.ConversionActionCategoryEnum_UNSPECIFIED
}

func (m *Segments) GetConversionActionName() *wrappers.StringValue {
	if m != nil {
		return m.ConversionActionName
	}
	return nil
}

func (m *Segments) GetConversionAdjustment() *wrappers.BoolValue {
	if m != nil {
		return m.ConversionAdjustment
	}
	return nil
}

func (m *Segments) GetConversionAttributionEventType() enums.ConversionAttributionEventTypeEnum_ConversionAttributionEventType {
	if m != nil {
		return m.ConversionAttributionEventType
	}
	return enums.ConversionAttributionEventTypeEnum_UNSPECIFIED
}

func (m *Segments) GetConversionLagBucket() enums.ConversionLagBucketEnum_ConversionLagBucket {
	if m != nil {
		return m.ConversionLagBucket
	}
	return enums.ConversionLagBucketEnum_UNSPECIFIED
}

func (m *Segments) GetConversionOrAdjustmentLagBucket() enums.ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket {
	if m != nil {
		return m.ConversionOrAdjustmentLagBucket
	}
	return enums.ConversionOrAdjustmentLagBucketEnum_UNSPECIFIED
}

func (m *Segments) GetDate() *wrappers.StringValue {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Segments) GetDayOfWeek() enums.DayOfWeekEnum_DayOfWeek {
	if m != nil {
		return m.DayOfWeek
	}
	return enums.DayOfWeekEnum_UNSPECIFIED
}

func (m *Segments) GetDevice() enums.DeviceEnum_Device {
	if m != nil {
		return m.Device
	}
	return enums.DeviceEnum_UNSPECIFIED
}

func (m *Segments) GetExternalConversionSource() enums.ExternalConversionSourceEnum_ExternalConversionSource {
	if m != nil {
		return m.ExternalConversionSource
	}
	return enums.ExternalConversionSourceEnum_UNSPECIFIED
}

func (m *Segments) GetGeoTargetCity() *wrappers.StringValue {
	if m != nil {
		return m.GeoTargetCity
	}
	return nil
}

func (m *Segments) GetGeoTargetMetro() *wrappers.StringValue {
	if m != nil {
		return m.GeoTargetMetro
	}
	return nil
}

func (m *Segments) GetGeoTargetRegion() *wrappers.StringValue {
	if m != nil {
		return m.GeoTargetRegion
	}
	return nil
}

func (m *Segments) GetHotelBookingWindowDays() *wrappers.Int64Value {
	if m != nil {
		return m.HotelBookingWindowDays
	}
	return nil
}

func (m *Segments) GetHotelCenterId() *wrappers.Int64Value {
	if m != nil {
		return m.HotelCenterId
	}
	return nil
}

func (m *Segments) GetHotelCheckInDate() *wrappers.StringValue {
	if m != nil {
		return m.HotelCheckInDate
	}
	return nil
}

func (m *Segments) GetHotelCheckInDayOfWeek() enums.DayOfWeekEnum_DayOfWeek {
	if m != nil {
		return m.HotelCheckInDayOfWeek
	}
	return enums.DayOfWeekEnum_UNSPECIFIED
}

func (m *Segments) GetHotelCity() *wrappers.StringValue {
	if m != nil {
		return m.HotelCity
	}
	return nil
}

func (m *Segments) GetHotelClass() *wrappers.Int32Value {
	if m != nil {
		return m.HotelClass
	}
	return nil
}

func (m *Segments) GetHotelCountry() *wrappers.StringValue {
	if m != nil {
		return m.HotelCountry
	}
	return nil
}

func (m *Segments) GetHotelDateSelectionType() enums.HotelDateSelectionTypeEnum_HotelDateSelectionType {
	if m != nil {
		return m.HotelDateSelectionType
	}
	return enums.HotelDateSelectionTypeEnum_UNSPECIFIED
}

func (m *Segments) GetHotelLengthOfStay() *wrappers.Int32Value {
	if m != nil {
		return m.HotelLengthOfStay
	}
	return nil
}

func (m *Segments) GetHotelState() *wrappers.StringValue {
	if m != nil {
		return m.HotelState
	}
	return nil
}

func (m *Segments) GetHour() *wrappers.Int32Value {
	if m != nil {
		return m.Hour
	}
	return nil
}

func (m *Segments) GetInteractionOnThisExtension() *wrappers.BoolValue {
	if m != nil {
		return m.InteractionOnThisExtension
	}
	return nil
}

func (m *Segments) GetKeyword() *Keyword {
	if m != nil {
		return m.Keyword
	}
	return nil
}

func (m *Segments) GetMonth() *wrappers.StringValue {
	if m != nil {
		return m.Month
	}
	return nil
}

func (m *Segments) GetMonthOfYear() enums.MonthOfYearEnum_MonthOfYear {
	if m != nil {
		return m.MonthOfYear
	}
	return enums.MonthOfYearEnum_UNSPECIFIED
}

func (m *Segments) GetPartnerHotelId() *wrappers.StringValue {
	if m != nil {
		return m.PartnerHotelId
	}
	return nil
}

func (m *Segments) GetPlaceholderType() enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderType
	}
	return enums.PlaceholderTypeEnum_UNSPECIFIED
}

func (m *Segments) GetProductAggregatorId() *wrappers.UInt64Value {
	if m != nil {
		return m.ProductAggregatorId
	}
	return nil
}

func (m *Segments) GetProductBiddingCategoryLevel1() *wrappers.StringValue {
	if m != nil {
		return m.ProductBiddingCategoryLevel1
	}
	return nil
}

func (m *Segments) GetProductBiddingCategoryLevel2() *wrappers.StringValue {
	if m != nil {
		return m.ProductBiddingCategoryLevel2
	}
	return nil
}

func (m *Segments) GetProductBiddingCategoryLevel3() *wrappers.StringValue {
	if m != nil {
		return m.ProductBiddingCategoryLevel3
	}
	return nil
}

func (m *Segments) GetProductBiddingCategoryLevel4() *wrappers.StringValue {
	if m != nil {
		return m.ProductBiddingCategoryLevel4
	}
	return nil
}

func (m *Segments) GetProductBiddingCategoryLevel5() *wrappers.StringValue {
	if m != nil {
		return m.ProductBiddingCategoryLevel5
	}
	return nil
}

func (m *Segments) GetProductBrand() *wrappers.StringValue {
	if m != nil {
		return m.ProductBrand
	}
	return nil
}

func (m *Segments) GetProductChannel() enums.ProductChannelEnum_ProductChannel {
	if m != nil {
		return m.ProductChannel
	}
	return enums.ProductChannelEnum_UNSPECIFIED
}

func (m *Segments) GetProductChannelExclusivity() enums.ProductChannelExclusivityEnum_ProductChannelExclusivity {
	if m != nil {
		return m.ProductChannelExclusivity
	}
	return enums.ProductChannelExclusivityEnum_UNSPECIFIED
}

func (m *Segments) GetProductCondition() enums.ProductConditionEnum_ProductCondition {
	if m != nil {
		return m.ProductCondition
	}
	return enums.ProductConditionEnum_UNSPECIFIED
}

func (m *Segments) GetProductCountry() *wrappers.StringValue {
	if m != nil {
		return m.ProductCountry
	}
	return nil
}

func (m *Segments) GetProductCustomAttribute0() *wrappers.StringValue {
	if m != nil {
		return m.ProductCustomAttribute0
	}
	return nil
}

func (m *Segments) GetProductCustomAttribute1() *wrappers.StringValue {
	if m != nil {
		return m.ProductCustomAttribute1
	}
	return nil
}

func (m *Segments) GetProductCustomAttribute2() *wrappers.StringValue {
	if m != nil {
		return m.ProductCustomAttribute2
	}
	return nil
}

func (m *Segments) GetProductCustomAttribute3() *wrappers.StringValue {
	if m != nil {
		return m.ProductCustomAttribute3
	}
	return nil
}

func (m *Segments) GetProductCustomAttribute4() *wrappers.StringValue {
	if m != nil {
		return m.ProductCustomAttribute4
	}
	return nil
}

func (m *Segments) GetProductItemId() *wrappers.StringValue {
	if m != nil {
		return m.ProductItemId
	}
	return nil
}

func (m *Segments) GetProductLanguage() *wrappers.StringValue {
	if m != nil {
		return m.ProductLanguage
	}
	return nil
}

func (m *Segments) GetProductMerchantId() *wrappers.UInt64Value {
	if m != nil {
		return m.ProductMerchantId
	}
	return nil
}

func (m *Segments) GetProductStoreId() *wrappers.StringValue {
	if m != nil {
		return m.ProductStoreId
	}
	return nil
}

func (m *Segments) GetProductTitle() *wrappers.StringValue {
	if m != nil {
		return m.ProductTitle
	}
	return nil
}

func (m *Segments) GetProductTypeL1() *wrappers.StringValue {
	if m != nil {
		return m.ProductTypeL1
	}
	return nil
}

func (m *Segments) GetProductTypeL2() *wrappers.StringValue {
	if m != nil {
		return m.ProductTypeL2
	}
	return nil
}

func (m *Segments) GetProductTypeL3() *wrappers.StringValue {
	if m != nil {
		return m.ProductTypeL3
	}
	return nil
}

func (m *Segments) GetProductTypeL4() *wrappers.StringValue {
	if m != nil {
		return m.ProductTypeL4
	}
	return nil
}

func (m *Segments) GetProductTypeL5() *wrappers.StringValue {
	if m != nil {
		return m.ProductTypeL5
	}
	return nil
}

func (m *Segments) GetQuarter() *wrappers.StringValue {
	if m != nil {
		return m.Quarter
	}
	return nil
}

func (m *Segments) GetSearchTermMatchType() enums.SearchTermMatchTypeEnum_SearchTermMatchType {
	if m != nil {
		return m.SearchTermMatchType
	}
	return enums.SearchTermMatchTypeEnum_UNSPECIFIED
}

func (m *Segments) GetSlot() enums.SlotEnum_Slot {
	if m != nil {
		return m.Slot
	}
	return enums.SlotEnum_UNSPECIFIED
}

func (m *Segments) GetWeek() *wrappers.StringValue {
	if m != nil {
		return m.Week
	}
	return nil
}

func (m *Segments) GetYear() *wrappers.Int32Value {
	if m != nil {
		return m.Year
	}
	return nil
}

// A Keyword criterion segment.
type Keyword struct {
	// The AdGroupCriterion resource name.
	AdGroupCriterion *wrappers.StringValue `protobuf:"bytes,1,opt,name=ad_group_criterion,json=adGroupCriterion,proto3" json:"ad_group_criterion,omitempty"`
	// Keyword info.
	Info                 *KeywordInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Keyword) Reset()         { *m = Keyword{} }
func (m *Keyword) String() string { return proto.CompactTextString(m) }
func (*Keyword) ProtoMessage()    {}
func (*Keyword) Descriptor() ([]byte, []int) {
	return fileDescriptor_segments_74ce948b3f70e9be, []int{1}
}
func (m *Keyword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyword.Unmarshal(m, b)
}
func (m *Keyword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyword.Marshal(b, m, deterministic)
}
func (dst *Keyword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyword.Merge(dst, src)
}
func (m *Keyword) XXX_Size() int {
	return xxx_messageInfo_Keyword.Size(m)
}
func (m *Keyword) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyword.DiscardUnknown(m)
}

var xxx_messageInfo_Keyword proto.InternalMessageInfo

func (m *Keyword) GetAdGroupCriterion() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupCriterion
	}
	return nil
}

func (m *Keyword) GetInfo() *KeywordInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Segments)(nil), "google.ads.googleads.v1.common.Segments")
	proto.RegisterType((*Keyword)(nil), "google.ads.googleads.v1.common.Keyword")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/segments.proto", fileDescriptor_segments_74ce948b3f70e9be)
}

var fileDescriptor_segments_74ce948b3f70e9be = []byte{
	// 1853 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x99, 0xcd, 0x76, 0xdb, 0xc6,
	0x15, 0xc7, 0x0f, 0x1d, 0xc5, 0x8e, 0x47, 0x91, 0x25, 0x8d, 0x6c, 0x67, 0x2c, 0x3b, 0xb6, 0xa3,
	0xb4, 0x8d, 0x1b, 0x27, 0xa4, 0xf8, 0x21, 0xb7, 0xa1, 0xe3, 0x0f, 0x8a, 0x52, 0x65, 0x26, 0x72,
	0xa4, 0x03, 0x29, 0x4a, 0xdb, 0xe3, 0x16, 0x1d, 0x01, 0x23, 0x10, 0x15, 0x30, 0x43, 0x0f, 0x86,
	0x52, 0xb8, 0x69, 0x1f, 0xa0, 0xbb, 0x2e, 0xfc, 0x00, 0x5d, 0xb6, 0x9b, 0x3e, 0x47, 0x9f, 0xa4,
	0xa7, 0x0f, 0xd0, 0x75, 0xce, 0x7c, 0x00, 0x24, 0x45, 0x91, 0x18, 0x4a, 0x2b, 0x01, 0x83, 0xfb,
	0xfb, 0xdf, 0x8b, 0xcb, 0x8b, 0x3b, 0x1f, 0x02, 0x5f, 0x06, 0x8c, 0x05, 0x11, 0x29, 0x61, 0x3f,
	0x29, 0xe9, 0x4b, 0x79, 0x75, 0x52, 0x2e, 0x79, 0x2c, 0x8e, 0x19, 0x2d, 0x25, 0x24, 0x88, 0x09,
	0x15, 0x49, 0xb1, 0xc3, 0x99, 0x60, 0xf0, 0xbe, 0xb6, 0x29, 0x62, 0x3f, 0x29, 0x66, 0xe6, 0xc5,
	0x93, 0x72, 0x51, 0x9b, 0x2f, 0xe7, 0xc9, 0x79, 0x3c, 0x14, 0x84, 0x87, 0x58, 0xcb, 0x2d, 0x57,
	0xc7, 0x99, 0x13, 0xda, 0x8d, 0x93, 0x12, 0xf6, 0x5d, 0x4a, 0xc4, 0x29, 0xe3, 0xc7, 0xae, 0xe8,
	0x75, 0x88, 0x81, 0x8a, 0x93, 0x21, 0x2f, 0x0a, 0xbd, 0x21, 0xfb, 0xe7, 0x39, 0xf6, 0x8c, 0x9e,
	0x10, 0x9e, 0x84, 0x8c, 0xba, 0xd8, 0x13, 0xf2, 0x8f, 0x87, 0x05, 0x09, 0x18, 0xef, 0x19, 0x7e,
	0xd3, 0x9e, 0x17, 0x82, 0x87, 0x87, 0x5d, 0x25, 0x42, 0x4e, 0x08, 0x15, 0x83, 0x61, 0x7c, 0x65,
	0x2d, 0x13, 0xe1, 0xc0, 0x3d, 0xec, 0x7a, 0xc7, 0x44, 0x18, 0x74, 0xcb, 0x1a, 0x65, 0xdc, 0xc5,
	0xfe, 0x9f, 0xbb, 0x89, 0x90, 0x3f, 0xda, 0xa8, 0x50, 0x69, 0xb2, 0x90, 0x8f, 0x7b, 0x2e, 0x3b,
	0x72, 0x4f, 0x09, 0x39, 0x36, 0xc0, 0xe7, 0x39, 0x00, 0x39, 0x09, 0x3d, 0xcb, 0x3c, 0x93, 0x1f,
	0x05, 0xe1, 0x14, 0x47, 0xee, 0x40, 0xb8, 0x09, 0xeb, 0xf2, 0x8c, 0x7f, 0x36, 0x99, 0x6f, 0x33,
	0x41, 0x22, 0xd7, 0xc7, 0x82, 0xb8, 0x09, 0x89, 0x88, 0xfe, 0xa9, 0x06, 0xf2, 0x5b, 0x9e, 0x8c,
	0xc7, 0x8c, 0x8a, 0xb6, 0x7c, 0xbb, 0x1e, 0xc1, 0xdc, 0x20, 0xb5, 0xc9, 0x48, 0x27, 0xc2, 0x1e,
	0x69, 0xb3, 0xc8, 0x27, 0x7c, 0xd0, 0x51, 0x4e, 0xd1, 0x76, 0x38, 0xf3, 0xbb, 0x9e, 0x70, 0xbd,
	0x36, 0xa6, 0x94, 0x44, 0x06, 0x7a, 0x31, 0x15, 0xe4, 0x92, 0x1f, 0xbd, 0xa8, 0x9b, 0x84, 0x27,
	0xa1, 0x48, 0xab, 0x70, 0xcd, 0x52, 0x80, 0x51, 0x3f, 0x94, 0xa9, 0x31, 0x58, 0x7d, 0x32, 0x96,
	0x10, 0xcc, 0xbd, 0xb6, 0x2b, 0x08, 0x8f, 0xdd, 0x18, 0x0b, 0x79, 0xd9, 0x7f, 0xd1, 0x47, 0x39,
	0x6c, 0xc4, 0xd2, 0xba, 0x32, 0x6d, 0xa1, 0xa4, 0xee, 0x0e, 0xbb, 0x47, 0xa5, 0x53, 0x8e, 0x3b,
	0x1d, 0xc2, 0x4d, 0xdb, 0x58, 0xbe, 0x97, 0x2a, 0x75, 0xc2, 0x12, 0xa6, 0x94, 0x09, 0x2c, 0x43,
	0x34, 0x4f, 0x57, 0xfe, 0xff, 0x18, 0x7c, 0xb0, 0x67, 0xfa, 0x0c, 0x3c, 0x02, 0xf3, 0x67, 0x3e,
	0x7b, 0xf4, 0xde, 0xc3, 0xc2, 0xa3, 0x1b, 0x95, 0xe7, 0xc5, 0x71, 0xbd, 0x47, 0x85, 0x53, 0x6c,
	0xf8, 0xdf, 0x69, 0x68, 0xbf, 0xd7, 0x21, 0x9b, 0xb4, 0x1b, 0x0f, 0x8f, 0x38, 0x73, 0x78, 0xf0,
	0x16, 0x7e, 0x0f, 0x40, 0xbf, 0x53, 0xa0, 0x65, 0xe5, 0xe2, 0x49, 0x8e, 0x8b, 0xa6, 0x04, 0x32,
	0xf9, 0xec, 0xce, 0xb9, 0xee, 0xa5, 0x97, 0xb0, 0x05, 0x16, 0x47, 0x1a, 0x0a, 0xaa, 0x3d, 0x2c,
	0x3c, 0x9a, 0xad, 0xdc, 0x4b, 0xd5, 0xd3, 0x2c, 0x15, 0xf7, 0x04, 0x0f, 0x69, 0x70, 0x80, 0xa3,
	0x2e, 0x71, 0x16, 0xfa, 0x58, 0x43, 0x51, 0xf0, 0xef, 0x05, 0xb0, 0x3c, 0xbe, 0x39, 0xa1, 0x35,
	0x15, 0xf2, 0x7e, 0x5e, 0xc8, 0x67, 0x54, 0x9b, 0x06, 0xd7, 0x6f, 0x30, 0xe6, 0xa1, 0x83, 0xbc,
	0x31, 0x4f, 0xa0, 0x03, 0x6e, 0x8f, 0xc6, 0x44, 0x71, 0x4c, 0xd0, 0x13, 0x8b, 0x97, 0xbc, 0x79,
	0x56, 0xf7, 0x3b, 0x1c, 0x13, 0xb8, 0x03, 0x6e, 0x0d, 0x6a, 0x66, 0xfd, 0x0b, 0xdd, 0x55, 0x92,
	0xcb, 0x23, 0x92, 0xeb, 0x8c, 0x45, 0xa3, 0x82, 0x19, 0x07, 0xff, 0x55, 0x00, 0x9f, 0xe4, 0xb6,
	0x65, 0x74, 0x45, 0x25, 0xf0, 0x4f, 0xf6, 0x09, 0xec, 0xcb, 0x6c, 0x4a, 0x95, 0x7e, 0x21, 0x4c,
	0x34, 0x71, 0xee, 0x7b, 0x13, 0x9f, 0xc3, 0xbf, 0x0c, 0xbd, 0x7e, 0xbf, 0x67, 0xa3, 0x8a, 0x0a,
	0xf0, 0x1b, 0xeb, 0x00, 0xb7, 0x71, 0xb0, 0xae, 0xc8, 0x33, 0x51, 0x65, 0xe3, 0xce, 0x92, 0x37,
	0x3a, 0x08, 0xff, 0x5d, 0x00, 0x9f, 0x5a, 0x4c, 0x21, 0xa8, 0xaa, 0xc2, 0x39, 0xb4, 0x0e, 0x67,
	0x87, 0xf7, 0x7f, 0x92, 0x71, 0xa1, 0x9d, 0x6b, 0xe3, 0x3c, 0xf0, 0x26, 0x1b, 0xc0, 0x55, 0x30,
	0x23, 0x27, 0x02, 0x34, 0x63, 0x51, 0x73, 0xca, 0x12, 0x1e, 0x80, 0xd9, 0x81, 0xd9, 0x0d, 0xbd,
	0x6f, 0xf5, 0xbd, 0x6f, 0xe0, 0xde, 0xce, 0xd1, 0x0f, 0x84, 0x1c, 0xab, 0xa8, 0xb3, 0x3b, 0xe7,
	0xba, 0x9f, 0x5e, 0xc2, 0x57, 0xe0, 0xaa, 0x9e, 0x04, 0x51, 0x41, 0x49, 0xae, 0xe6, 0x49, 0x2a,
	0x63, 0xad, 0xa7, 0x2e, 0x1d, 0xc3, 0xab, 0xcf, 0x7d, 0xfc, 0x1c, 0x89, 0x7e, 0x65, 0xf5, 0xb9,
	0x6f, 0x1a, 0x81, 0x7e, 0x86, 0xf7, 0x14, 0xae, 0x1c, 0x8e, 0x7b, 0xe8, 0x20, 0x32, 0xe6, 0x09,
	0xdc, 0x00, 0xf3, 0x01, 0x61, 0xae, 0xc0, 0x3c, 0x20, 0xc2, 0xf5, 0x42, 0xd1, 0x43, 0xcf, 0x2d,
	0x72, 0x3e, 0x17, 0x10, 0xb6, 0xaf, 0x98, 0x66, 0x28, 0x7a, 0xf0, 0x37, 0x60, 0x61, 0x40, 0x25,
	0x26, 0x82, 0x33, 0xf4, 0xc2, 0x42, 0xe6, 0x46, 0x26, 0xf3, 0x5a, 0x32, 0xf0, 0x15, 0x58, 0x1c,
	0xd0, 0xe1, 0x24, 0x90, 0xcd, 0xf5, 0xa5, 0x85, 0xd0, 0x7c, 0x26, 0xe4, 0x28, 0x08, 0x1e, 0x80,
	0x3b, 0x7a, 0x3d, 0x71, 0xc8, 0xd8, 0x71, 0x48, 0x03, 0xf7, 0x34, 0xa4, 0x3e, 0x3b, 0x75, 0x7d,
	0xdc, 0x4b, 0xd0, 0x55, 0xa5, 0x78, 0x77, 0x44, 0xb1, 0x45, 0xc5, 0x93, 0x9a, 0x16, 0xbc, 0xad,
	0xe8, 0x75, 0x0d, 0xff, 0xa0, 0xd8, 0x0d, 0xdc, 0x4b, 0x60, 0x13, 0xcc, 0x6b, 0x5d, 0x8f, 0x50,
	0x41, 0xb8, 0x1b, 0xfa, 0xe8, 0x5a, 0xbe, 0xda, 0x9c, 0x62, 0x9a, 0x0a, 0x69, 0xf9, 0xf0, 0x5b,
	0xb0, 0x64, 0x44, 0xda, 0xc4, 0x3b, 0x76, 0x43, 0xaa, 0x56, 0x3d, 0xe8, 0x03, 0x9b, 0x59, 0x44,
	0x2b, 0x49, 0xae, 0x45, 0x37, 0x64, 0xe1, 0xbf, 0x05, 0xcb, 0x23, 0x62, 0xfd, 0xef, 0xe0, 0xfa,
	0xa5, 0xbe, 0x83, 0x5b, 0xc3, 0xde, 0xd2, 0x6f, 0xe2, 0x29, 0x00, 0xc6, 0xa5, 0xac, 0x17, 0x60,
	0x11, 0xf6, 0x75, 0x2d, 0x24, 0x6b, 0xe5, 0x6b, 0x30, 0x6b, 0xe0, 0x08, 0x27, 0x09, 0x9a, 0x1d,
	0x9f, 0xbd, 0x6a, 0x45, 0xc3, 0xda, 0x59, 0x53, 0x9a, 0xc3, 0x06, 0x98, 0x33, 0x34, 0xeb, 0x52,
	0xc1, 0x7b, 0xe8, 0x43, 0x0b, 0xef, 0x1f, 0x6a, 0x01, 0x4d, 0xc0, 0xbf, 0x15, 0xd2, 0xda, 0x38,
	0x67, 0xad, 0x89, 0xe6, 0x54, 0xc2, 0x76, 0x73, 0x12, 0xf6, 0x4a, 0xf2, 0x32, 0xfd, 0x7b, 0x29,
	0x9d, 0x4d, 0x16, 0xe7, 0x3f, 0x32, 0x05, 0x35, 0x32, 0x0e, 0xb7, 0xc1, 0x4d, 0x1d, 0x4c, 0x44,
	0x68, 0xa0, 0x17, 0xb0, 0x89, 0xc0, 0x3d, 0x74, 0x23, 0x3f, 0x2f, 0x8b, 0x0a, 0xdc, 0x56, 0xdc,
	0xce, 0xd1, 0x9e, 0xc0, 0x3d, 0xf8, 0x2c, 0x4d, 0x6e, 0x22, 0x64, 0x45, 0xcd, 0x5b, 0x24, 0x47,
	0x67, 0x77, 0x4f, 0xda, 0xc3, 0x12, 0x98, 0x69, 0xb3, 0x2e, 0x47, 0x0b, 0xf9, 0xce, 0x95, 0x21,
	0xfc, 0x03, 0xf8, 0x38, 0x94, 0x45, 0x6d, 0xd6, 0x09, 0x32, 0x89, 0xed, 0x30, 0x71, 0x65, 0xbb,
	0xa1, 0xb2, 0xcb, 0xa0, 0x72, 0xee, 0x0c, 0xbf, 0x3c, 0x20, 0xb0, 0x43, 0xf7, 0xdb, 0x61, 0xb2,
	0x99, 0xd2, 0xb0, 0x01, 0xae, 0x1d, 0x93, 0xde, 0x29, 0xe3, 0x3e, 0x7a, 0xa6, 0x84, 0x3e, 0x2b,
	0x4e, 0xde, 0x9f, 0x16, 0xbf, 0xd5, 0xe6, 0x4e, 0xca, 0xc1, 0x0a, 0x78, 0x5f, 0xed, 0x0c, 0xd0,
	0xa2, 0x45, 0x2e, 0xb4, 0x29, 0xfc, 0x23, 0x98, 0x1b, 0xda, 0x4d, 0x20, 0xa8, 0x8a, 0xa2, 0x9e,
	0x53, 0x14, 0xaf, 0x25, 0xb3, 0x73, 0xf4, 0x3b, 0x82, 0xb9, 0xaa, 0x84, 0x81, 0x7b, 0x67, 0x36,
	0xee, 0xdf, 0xc8, 0x76, 0xd9, 0xc1, 0x5c, 0x50, 0xc2, 0x5d, 0xfd, 0x6b, 0x85, 0x3e, 0x5a, 0xb2,
	0x69, 0x97, 0x86, 0x52, 0x25, 0xd6, 0xf2, 0x61, 0x0c, 0x16, 0xce, 0x6e, 0x61, 0xd0, 0x4d, 0x15,
	0xea, 0x7a, 0x4e, 0xa8, 0xbb, 0x7d, 0x2c, 0x2b, 0xdc, 0x33, 0x63, 0xce, 0x7c, 0x67, 0x78, 0x00,
	0xee, 0x82, 0x5b, 0xe9, 0x2e, 0x04, 0x07, 0x01, 0x27, 0x01, 0x16, 0x4c, 0x75, 0xc0, 0x7b, 0x63,
	0x62, 0xff, 0x7e, 0xa0, 0x05, 0x2e, 0x19, 0xb4, 0x91, 0x91, 0x2d, 0x1f, 0x7a, 0xe0, 0x41, 0xaa,
	0x78, 0x18, 0xfa, 0xbe, 0xec, 0xd3, 0xe9, 0xf2, 0xd7, 0x8d, 0xc8, 0x09, 0x89, 0xca, 0xe8, 0xd7,
	0x16, 0x79, 0xb9, 0x67, 0x44, 0xd6, 0xb5, 0x46, 0xba, 0x96, 0xdd, 0x56, 0x0a, 0xf9, 0x4e, 0x2a,
	0xe8, 0xab, 0x4b, 0x3a, 0xa9, 0xe4, 0x3b, 0xa9, 0xa2, 0xfa, 0x25, 0x9d, 0x54, 0xf3, 0x9d, 0xd4,
	0xd0, 0xd3, 0x4b, 0x3a, 0xa9, 0xe5, 0x3b, 0x59, 0x43, 0x5f, 0x5f, 0xd2, 0xc9, 0x9a, 0x6c, 0xe3,
	0x99, 0x13, 0x8e, 0xa9, 0x8f, 0x3e, 0xb6, 0x69, 0xe3, 0xa9, 0xa4, 0x24, 0x60, 0x08, 0xe6, 0xcf,
	0x6c, 0xaa, 0xd1, 0x7d, 0x55, 0xfb, 0x2f, 0xf3, 0x6a, 0x5f, 0x53, 0x4d, 0x0d, 0xe9, 0xd2, 0x1f,
	0x1a, 0x72, 0x6e, 0x74, 0x86, 0xee, 0xe1, 0xbb, 0x02, 0xb8, 0x3b, 0x61, 0x03, 0x8f, 0x1e, 0x28,
	0xbf, 0x07, 0xd3, 0xf9, 0xed, 0xf3, 0xe7, 0x84, 0x30, 0xf0, 0xd4, 0xb9, 0xd3, 0x19, 0xf7, 0x08,
	0xbe, 0x05, 0x8b, 0x23, 0xe7, 0x02, 0xe8, 0xa1, 0x8a, 0x66, 0xc3, 0x32, 0x9a, 0x14, 0x1b, 0x0a,
	0x22, 0x1d, 0x74, 0x16, 0x3a, 0x67, 0x46, 0xe0, 0xe6, 0x40, 0xda, 0xcd, 0x14, 0xfc, 0x89, 0x55,
	0xeb, 0x4a, 0x85, 0xf4, 0x24, 0xfc, 0x5b, 0x70, 0x27, 0x93, 0xe9, 0x26, 0x82, 0xc5, 0xd9, 0x26,
	0x8e, 0xac, 0xa2, 0x15, 0x0b, 0xc1, 0x8f, 0x52, 0x41, 0x45, 0x37, 0x32, 0x78, 0x92, 0x72, 0x19,
	0x7d, 0x7a, 0x71, 0xe5, 0xf2, 0x24, 0xe5, 0x0a, 0xfa, 0xd9, 0xc5, 0x95, 0x2b, 0x93, 0x94, 0xab,
	0xe8, 0xe7, 0x17, 0x57, 0xae, 0x4e, 0x52, 0xae, 0xa1, 0x5f, 0x5c, 0x5c, 0xb9, 0x26, 0x77, 0x0e,
	0xa9, 0x72, 0x28, 0x48, 0x2c, 0xe7, 0x81, 0xcf, 0x6c, 0x76, 0x0e, 0x06, 0x6a, 0x09, 0x12, 0xb7,
	0x7c, 0xb8, 0x05, 0xd2, 0x12, 0x73, 0x23, 0x4c, 0x83, 0x2e, 0x0e, 0x08, 0x7a, 0x64, 0xb3, 0xe0,
	0x37, 0xd4, 0xb6, 0x81, 0xe0, 0x36, 0x48, 0x67, 0x18, 0x37, 0x26, 0x5c, 0x7e, 0xa5, 0x42, 0x86,
	0xf4, 0x4b, 0x8b, 0xa9, 0x29, 0xfd, 0x86, 0x5e, 0x1b, 0xae, 0xe5, 0xab, 0x19, 0xda, 0xa8, 0x25,
	0x82, 0x71, 0x22, 0xa5, 0x3e, 0x9f, 0xa2, 0xcc, 0xf7, 0x24, 0xd4, 0xf2, 0x07, 0xfb, 0x9c, 0x08,
	0x45, 0x44, 0xd0, 0xe3, 0x29, 0xfa, 0xdc, 0xbe, 0x24, 0x06, 0xf3, 0x2c, 0x27, 0x78, 0x37, 0x2a,
	0xa3, 0x2f, 0xa6, 0xc8, 0xb3, 0x9c, 0xb9, 0xb7, 0xcb, 0xa3, 0x2a, 0x15, 0xf4, 0xe5, 0xb4, 0x2a,
	0x95, 0x51, 0x95, 0x2a, 0x2a, 0x4e, 0xab, 0x52, 0x1d, 0x55, 0xa9, 0xa1, 0xd2, 0xb4, 0x2a, 0xb5,
	0x51, 0x95, 0x35, 0xb4, 0x3a, 0xad, 0xca, 0x1a, 0x7c, 0x02, 0xae, 0xbd, 0xed, 0x62, 0x2e, 0x08,
	0x47, 0xb7, 0x2c, 0xe8, 0xd4, 0x18, 0xfe, 0x15, 0xdc, 0x3e, 0xff, 0x68, 0x15, 0xdd, 0xb6, 0x3a,
	0xd4, 0xd9, 0x53, 0xf0, 0x3e, 0xe1, 0xf1, 0x6b, 0x89, 0x66, 0x8b, 0xb0, 0x73, 0xc6, 0x9d, 0xa5,
	0x64, 0x74, 0x10, 0xbe, 0x04, 0x33, 0x49, 0xc4, 0x04, 0xfa, 0x48, 0xb9, 0xfb, 0x22, 0xcf, 0x5d,
	0xc4, 0xf4, 0xc9, 0x8c, 0xbc, 0x70, 0x14, 0x09, 0x57, 0xc1, 0x8c, 0xda, 0x22, 0x22, 0x9b, 0x33,
	0x16, 0x69, 0x29, 0xb7, 0x07, 0x6a, 0x39, 0x7c, 0xc7, 0x62, 0x7b, 0x20, 0x0d, 0x57, 0xde, 0x15,
	0xc0, 0x35, 0xb3, 0x22, 0x87, 0xdf, 0x00, 0x88, 0x7d, 0x37, 0xe0, 0xac, 0xdb, 0x71, 0xcd, 0x7f,
	0x89, 0x18, 0x55, 0x87, 0x2a, 0xb9, 0x7b, 0x5e, 0xec, 0x6f, 0x49, 0xac, 0x99, 0x52, 0xf0, 0x05,
	0x98, 0x09, 0xe9, 0x11, 0x53, 0x27, 0x7c, 0xb3, 0x95, 0xc7, 0x96, 0x9b, 0x82, 0x16, 0x3d, 0x62,
	0x8e, 0x02, 0xd7, 0xff, 0x5b, 0x00, 0x2b, 0x1e, 0x8b, 0x73, 0xc0, 0xf5, 0xb9, 0xf4, 0xd4, 0x7a,
	0x57, 0xc6, 0xb5, 0x5b, 0xf8, 0xfd, 0x86, 0x01, 0x02, 0x26, 0xbb, 0x55, 0x91, 0xf1, 0xa0, 0x14,
	0x10, 0xaa, 0xa2, 0x4e, 0x4f, 0xd0, 0x3b, 0x61, 0x32, 0xee, 0xbf, 0x63, 0x4f, 0xf5, 0x9f, 0x7f,
	0x5c, 0x79, 0x6f, 0xab, 0xd1, 0xf8, 0xe7, 0x95, 0xfb, 0x5b, 0x5a, 0xac, 0xe1, 0x27, 0x45, 0x7d,
	0x29, 0xaf, 0x0e, 0xca, 0xc5, 0xa6, 0x32, 0xfb, 0x4f, 0x6a, 0xf0, 0xa6, 0xe1, 0x27, 0x6f, 0x32,
	0x83, 0x37, 0x07, 0xe5, 0x37, 0xda, 0xe0, 0x7f, 0x57, 0x56, 0xf4, 0x68, 0xbd, 0xde, 0xf0, 0x93,
	0x7a, 0x3d, 0x33, 0xa9, 0xd7, 0x0f, 0xca, 0xf5, 0xba, 0x36, 0x3a, 0xbc, 0xaa, 0xa2, 0xab, 0xfe,
	0x14, 0x00, 0x00, 0xff, 0xff, 0x55, 0xaa, 0xe8, 0x82, 0x09, 0x1c, 0x00, 0x00,
}
