package JYBaseDB

//This file is generate by scripts,don't edit it
//

import (
	"changit.cn/contra/bot/db"
)

func LoadBaseData() {
	SysAdCache.LoadAll()
	SysAdPositionCache.LoadAll()
	SysAdminCache.LoadAll()
	SysAnnounceCache.LoadAll()
	SysAnnounceViewCache.LoadAll()
	SysAppVersionCache.LoadAll()
	SysAuthGroupCache.LoadAll()
	SysAuthGroupAccessCache.LoadAll()
	SysAuthRuleCache.LoadAll()
	SysBlacklistCache.LoadAll()
	SysConfigCache.LoadAll()
	SysJpushDeviceCache.LoadAll()
	SysJpushDeviceLogCache.LoadAll()
	SysLinksCache.LoadAll()
	SysPushConfCache.LoadAll()
	SysPushSendContentCache.LoadAll()
	SysSensitiveWordsCache.LoadAll()
	SysSmsConfCache.LoadAll()
	SysSmsTplCache.LoadAll()
	SysZoneCache.LoadAll()
	db.BaseDataCaches["SysAd"] = SysAdCache
	db.BaseDataCaches["SysAdPosition"] = SysAdPositionCache
	db.BaseDataCaches["SysAdmin"] = SysAdminCache
	db.BaseDataCaches["SysAnnounce"] = SysAnnounceCache
	db.BaseDataCaches["SysAnnounceView"] = SysAnnounceViewCache
	db.BaseDataCaches["SysAppVersion"] = SysAppVersionCache
	db.BaseDataCaches["SysAuthGroup"] = SysAuthGroupCache
	db.BaseDataCaches["SysAuthGroupAccess"] = SysAuthGroupAccessCache
	db.BaseDataCaches["SysAuthRule"] = SysAuthRuleCache
	db.BaseDataCaches["SysBlacklist"] = SysBlacklistCache
	db.BaseDataCaches["SysConfig"] = SysConfigCache
	db.BaseDataCaches["SysJpushDevice"] = SysJpushDeviceCache
	db.BaseDataCaches["SysJpushDeviceLog"] = SysJpushDeviceLogCache
	db.BaseDataCaches["SysLinks"] = SysLinksCache
	db.BaseDataCaches["SysPushConf"] = SysPushConfCache
	db.BaseDataCaches["SysPushSendContent"] = SysPushSendContentCache
	db.BaseDataCaches["SysSensitiveWords"] = SysSensitiveWordsCache
	db.BaseDataCaches["SysSmsConf"] = SysSmsConfCache
	db.BaseDataCaches["SysSmsTpl"] = SysSmsTplCache
	db.BaseDataCaches["SysZone"] = SysZoneCache
}
