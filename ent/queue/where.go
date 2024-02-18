// Code generated by ent, DO NOT EDIT.

package queue

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/internal/utils"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldLTE(FieldID, id))
}

// LiveArchive applies equality check predicate on the "live_archive" field. It's identical to LiveArchiveEQ.
func LiveArchive(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldLiveArchive, v))
}

// OnHold applies equality check predicate on the "on_hold" field. It's identical to OnHoldEQ.
func OnHold(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldOnHold, v))
}

// VideoProcessing applies equality check predicate on the "video_processing" field. It's identical to VideoProcessingEQ.
func VideoProcessing(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldVideoProcessing, v))
}

// ChatProcessing applies equality check predicate on the "chat_processing" field. It's identical to ChatProcessingEQ.
func ChatProcessing(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldChatProcessing, v))
}

// Processing applies equality check predicate on the "processing" field. It's identical to ProcessingEQ.
func Processing(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldProcessing, v))
}

// ChatStart applies equality check predicate on the "chat_start" field. It's identical to ChatStartEQ.
func ChatStart(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldChatStart, v))
}

// RenderChat applies equality check predicate on the "render_chat" field. It's identical to RenderChatEQ.
func RenderChat(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldRenderChat, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldCreatedAt, v))
}

// LiveArchiveEQ applies the EQ predicate on the "live_archive" field.
func LiveArchiveEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldLiveArchive, v))
}

// LiveArchiveNEQ applies the NEQ predicate on the "live_archive" field.
func LiveArchiveNEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldLiveArchive, v))
}

// OnHoldEQ applies the EQ predicate on the "on_hold" field.
func OnHoldEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldOnHold, v))
}

// OnHoldNEQ applies the NEQ predicate on the "on_hold" field.
func OnHoldNEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldOnHold, v))
}

// VideoProcessingEQ applies the EQ predicate on the "video_processing" field.
func VideoProcessingEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldVideoProcessing, v))
}

// VideoProcessingNEQ applies the NEQ predicate on the "video_processing" field.
func VideoProcessingNEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldVideoProcessing, v))
}

// ChatProcessingEQ applies the EQ predicate on the "chat_processing" field.
func ChatProcessingEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldChatProcessing, v))
}

// ChatProcessingNEQ applies the NEQ predicate on the "chat_processing" field.
func ChatProcessingNEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldChatProcessing, v))
}

// ProcessingEQ applies the EQ predicate on the "processing" field.
func ProcessingEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldProcessing, v))
}

// ProcessingNEQ applies the NEQ predicate on the "processing" field.
func ProcessingNEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldProcessing, v))
}

// TaskVodCreateFolderEQ applies the EQ predicate on the "task_vod_create_folder" field.
func TaskVodCreateFolderEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldEQ(FieldTaskVodCreateFolder, vc))
}

// TaskVodCreateFolderNEQ applies the NEQ predicate on the "task_vod_create_folder" field.
func TaskVodCreateFolderNEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldNEQ(FieldTaskVodCreateFolder, vc))
}

// TaskVodCreateFolderIn applies the In predicate on the "task_vod_create_folder" field.
func TaskVodCreateFolderIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldIn(FieldTaskVodCreateFolder, v...))
}

// TaskVodCreateFolderNotIn applies the NotIn predicate on the "task_vod_create_folder" field.
func TaskVodCreateFolderNotIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldNotIn(FieldTaskVodCreateFolder, v...))
}

// TaskVodCreateFolderIsNil applies the IsNil predicate on the "task_vod_create_folder" field.
func TaskVodCreateFolderIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldTaskVodCreateFolder))
}

// TaskVodCreateFolderNotNil applies the NotNil predicate on the "task_vod_create_folder" field.
func TaskVodCreateFolderNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldTaskVodCreateFolder))
}

// TaskVodDownloadThumbnailEQ applies the EQ predicate on the "task_vod_download_thumbnail" field.
func TaskVodDownloadThumbnailEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldEQ(FieldTaskVodDownloadThumbnail, vc))
}

// TaskVodDownloadThumbnailNEQ applies the NEQ predicate on the "task_vod_download_thumbnail" field.
func TaskVodDownloadThumbnailNEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldNEQ(FieldTaskVodDownloadThumbnail, vc))
}

// TaskVodDownloadThumbnailIn applies the In predicate on the "task_vod_download_thumbnail" field.
func TaskVodDownloadThumbnailIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldIn(FieldTaskVodDownloadThumbnail, v...))
}

// TaskVodDownloadThumbnailNotIn applies the NotIn predicate on the "task_vod_download_thumbnail" field.
func TaskVodDownloadThumbnailNotIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldNotIn(FieldTaskVodDownloadThumbnail, v...))
}

// TaskVodDownloadThumbnailIsNil applies the IsNil predicate on the "task_vod_download_thumbnail" field.
func TaskVodDownloadThumbnailIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldTaskVodDownloadThumbnail))
}

// TaskVodDownloadThumbnailNotNil applies the NotNil predicate on the "task_vod_download_thumbnail" field.
func TaskVodDownloadThumbnailNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldTaskVodDownloadThumbnail))
}

// TaskVodSaveInfoEQ applies the EQ predicate on the "task_vod_save_info" field.
func TaskVodSaveInfoEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldEQ(FieldTaskVodSaveInfo, vc))
}

// TaskVodSaveInfoNEQ applies the NEQ predicate on the "task_vod_save_info" field.
func TaskVodSaveInfoNEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldNEQ(FieldTaskVodSaveInfo, vc))
}

// TaskVodSaveInfoIn applies the In predicate on the "task_vod_save_info" field.
func TaskVodSaveInfoIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldIn(FieldTaskVodSaveInfo, v...))
}

// TaskVodSaveInfoNotIn applies the NotIn predicate on the "task_vod_save_info" field.
func TaskVodSaveInfoNotIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldNotIn(FieldTaskVodSaveInfo, v...))
}

// TaskVodSaveInfoIsNil applies the IsNil predicate on the "task_vod_save_info" field.
func TaskVodSaveInfoIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldTaskVodSaveInfo))
}

// TaskVodSaveInfoNotNil applies the NotNil predicate on the "task_vod_save_info" field.
func TaskVodSaveInfoNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldTaskVodSaveInfo))
}

// TaskVideoDownloadEQ applies the EQ predicate on the "task_video_download" field.
func TaskVideoDownloadEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldEQ(FieldTaskVideoDownload, vc))
}

// TaskVideoDownloadNEQ applies the NEQ predicate on the "task_video_download" field.
func TaskVideoDownloadNEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldNEQ(FieldTaskVideoDownload, vc))
}

// TaskVideoDownloadIn applies the In predicate on the "task_video_download" field.
func TaskVideoDownloadIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldIn(FieldTaskVideoDownload, v...))
}

// TaskVideoDownloadNotIn applies the NotIn predicate on the "task_video_download" field.
func TaskVideoDownloadNotIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldNotIn(FieldTaskVideoDownload, v...))
}

// TaskVideoDownloadIsNil applies the IsNil predicate on the "task_video_download" field.
func TaskVideoDownloadIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldTaskVideoDownload))
}

// TaskVideoDownloadNotNil applies the NotNil predicate on the "task_video_download" field.
func TaskVideoDownloadNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldTaskVideoDownload))
}

// TaskVideoConvertEQ applies the EQ predicate on the "task_video_convert" field.
func TaskVideoConvertEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldEQ(FieldTaskVideoConvert, vc))
}

// TaskVideoConvertNEQ applies the NEQ predicate on the "task_video_convert" field.
func TaskVideoConvertNEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldNEQ(FieldTaskVideoConvert, vc))
}

// TaskVideoConvertIn applies the In predicate on the "task_video_convert" field.
func TaskVideoConvertIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldIn(FieldTaskVideoConvert, v...))
}

// TaskVideoConvertNotIn applies the NotIn predicate on the "task_video_convert" field.
func TaskVideoConvertNotIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldNotIn(FieldTaskVideoConvert, v...))
}

// TaskVideoConvertIsNil applies the IsNil predicate on the "task_video_convert" field.
func TaskVideoConvertIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldTaskVideoConvert))
}

// TaskVideoConvertNotNil applies the NotNil predicate on the "task_video_convert" field.
func TaskVideoConvertNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldTaskVideoConvert))
}

// TaskVideoMoveEQ applies the EQ predicate on the "task_video_move" field.
func TaskVideoMoveEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldEQ(FieldTaskVideoMove, vc))
}

// TaskVideoMoveNEQ applies the NEQ predicate on the "task_video_move" field.
func TaskVideoMoveNEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldNEQ(FieldTaskVideoMove, vc))
}

// TaskVideoMoveIn applies the In predicate on the "task_video_move" field.
func TaskVideoMoveIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldIn(FieldTaskVideoMove, v...))
}

// TaskVideoMoveNotIn applies the NotIn predicate on the "task_video_move" field.
func TaskVideoMoveNotIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldNotIn(FieldTaskVideoMove, v...))
}

// TaskVideoMoveIsNil applies the IsNil predicate on the "task_video_move" field.
func TaskVideoMoveIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldTaskVideoMove))
}

// TaskVideoMoveNotNil applies the NotNil predicate on the "task_video_move" field.
func TaskVideoMoveNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldTaskVideoMove))
}

// TaskChatDownloadEQ applies the EQ predicate on the "task_chat_download" field.
func TaskChatDownloadEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldEQ(FieldTaskChatDownload, vc))
}

// TaskChatDownloadNEQ applies the NEQ predicate on the "task_chat_download" field.
func TaskChatDownloadNEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldNEQ(FieldTaskChatDownload, vc))
}

// TaskChatDownloadIn applies the In predicate on the "task_chat_download" field.
func TaskChatDownloadIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldIn(FieldTaskChatDownload, v...))
}

// TaskChatDownloadNotIn applies the NotIn predicate on the "task_chat_download" field.
func TaskChatDownloadNotIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldNotIn(FieldTaskChatDownload, v...))
}

// TaskChatDownloadIsNil applies the IsNil predicate on the "task_chat_download" field.
func TaskChatDownloadIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldTaskChatDownload))
}

// TaskChatDownloadNotNil applies the NotNil predicate on the "task_chat_download" field.
func TaskChatDownloadNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldTaskChatDownload))
}

// TaskChatConvertEQ applies the EQ predicate on the "task_chat_convert" field.
func TaskChatConvertEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldEQ(FieldTaskChatConvert, vc))
}

// TaskChatConvertNEQ applies the NEQ predicate on the "task_chat_convert" field.
func TaskChatConvertNEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldNEQ(FieldTaskChatConvert, vc))
}

// TaskChatConvertIn applies the In predicate on the "task_chat_convert" field.
func TaskChatConvertIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldIn(FieldTaskChatConvert, v...))
}

// TaskChatConvertNotIn applies the NotIn predicate on the "task_chat_convert" field.
func TaskChatConvertNotIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldNotIn(FieldTaskChatConvert, v...))
}

// TaskChatConvertIsNil applies the IsNil predicate on the "task_chat_convert" field.
func TaskChatConvertIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldTaskChatConvert))
}

// TaskChatConvertNotNil applies the NotNil predicate on the "task_chat_convert" field.
func TaskChatConvertNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldTaskChatConvert))
}

// TaskChatRenderEQ applies the EQ predicate on the "task_chat_render" field.
func TaskChatRenderEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldEQ(FieldTaskChatRender, vc))
}

// TaskChatRenderNEQ applies the NEQ predicate on the "task_chat_render" field.
func TaskChatRenderNEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldNEQ(FieldTaskChatRender, vc))
}

// TaskChatRenderIn applies the In predicate on the "task_chat_render" field.
func TaskChatRenderIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldIn(FieldTaskChatRender, v...))
}

// TaskChatRenderNotIn applies the NotIn predicate on the "task_chat_render" field.
func TaskChatRenderNotIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldNotIn(FieldTaskChatRender, v...))
}

// TaskChatRenderIsNil applies the IsNil predicate on the "task_chat_render" field.
func TaskChatRenderIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldTaskChatRender))
}

// TaskChatRenderNotNil applies the NotNil predicate on the "task_chat_render" field.
func TaskChatRenderNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldTaskChatRender))
}

// TaskChatMoveEQ applies the EQ predicate on the "task_chat_move" field.
func TaskChatMoveEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldEQ(FieldTaskChatMove, vc))
}

// TaskChatMoveNEQ applies the NEQ predicate on the "task_chat_move" field.
func TaskChatMoveNEQ(v utils.TaskStatus) predicate.Queue {
	vc := v
	return predicate.Queue(sql.FieldNEQ(FieldTaskChatMove, vc))
}

// TaskChatMoveIn applies the In predicate on the "task_chat_move" field.
func TaskChatMoveIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldIn(FieldTaskChatMove, v...))
}

// TaskChatMoveNotIn applies the NotIn predicate on the "task_chat_move" field.
func TaskChatMoveNotIn(vs ...utils.TaskStatus) predicate.Queue {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Queue(sql.FieldNotIn(FieldTaskChatMove, v...))
}

// TaskChatMoveIsNil applies the IsNil predicate on the "task_chat_move" field.
func TaskChatMoveIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldTaskChatMove))
}

// TaskChatMoveNotNil applies the NotNil predicate on the "task_chat_move" field.
func TaskChatMoveNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldTaskChatMove))
}

// ChatStartEQ applies the EQ predicate on the "chat_start" field.
func ChatStartEQ(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldChatStart, v))
}

// ChatStartNEQ applies the NEQ predicate on the "chat_start" field.
func ChatStartNEQ(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldChatStart, v))
}

// ChatStartIn applies the In predicate on the "chat_start" field.
func ChatStartIn(vs ...time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldIn(FieldChatStart, vs...))
}

// ChatStartNotIn applies the NotIn predicate on the "chat_start" field.
func ChatStartNotIn(vs ...time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldNotIn(FieldChatStart, vs...))
}

// ChatStartGT applies the GT predicate on the "chat_start" field.
func ChatStartGT(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldGT(FieldChatStart, v))
}

// ChatStartGTE applies the GTE predicate on the "chat_start" field.
func ChatStartGTE(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldGTE(FieldChatStart, v))
}

// ChatStartLT applies the LT predicate on the "chat_start" field.
func ChatStartLT(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldLT(FieldChatStart, v))
}

// ChatStartLTE applies the LTE predicate on the "chat_start" field.
func ChatStartLTE(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldLTE(FieldChatStart, v))
}

// ChatStartIsNil applies the IsNil predicate on the "chat_start" field.
func ChatStartIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldChatStart))
}

// ChatStartNotNil applies the NotNil predicate on the "chat_start" field.
func ChatStartNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldChatStart))
}

// RenderChatEQ applies the EQ predicate on the "render_chat" field.
func RenderChatEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldRenderChat, v))
}

// RenderChatNEQ applies the NEQ predicate on the "render_chat" field.
func RenderChatNEQ(v bool) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldRenderChat, v))
}

// RenderChatIsNil applies the IsNil predicate on the "render_chat" field.
func RenderChatIsNil() predicate.Queue {
	return predicate.Queue(sql.FieldIsNull(FieldRenderChat))
}

// RenderChatNotNil applies the NotNil predicate on the "render_chat" field.
func RenderChatNotNil() predicate.Queue {
	return predicate.Queue(sql.FieldNotNull(FieldRenderChat))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Queue {
	return predicate.Queue(sql.FieldLTE(FieldCreatedAt, v))
}

// HasVod applies the HasEdge predicate on the "vod" edge.
func HasVod() predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, VodTable, VodColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVodWith applies the HasEdge predicate on the "vod" edge with a given conditions (other predicates).
func HasVodWith(preds ...predicate.Vod) predicate.Queue {
	return predicate.Queue(func(s *sql.Selector) {
		step := newVodStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Queue) predicate.Queue {
	return predicate.Queue(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Queue) predicate.Queue {
	return predicate.Queue(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Queue) predicate.Queue {
	return predicate.Queue(sql.NotPredicates(p))
}
