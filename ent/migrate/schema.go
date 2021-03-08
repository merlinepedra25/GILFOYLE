// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// MediaColumns holds the columns for the "media" table.
	MediaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "original_filename", Type: field.TypeString, Nullable: true, Size: 150, Default: ""},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"AwaitingUpload", "Processing", "Ready", "Errored"}},
		{Name: "message", Type: field.TypeString, Nullable: true, Size: 255, Default: ""},
		{Name: "playable", Type: field.TypeBool, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// MediaTable holds the schema information for the "media" table.
	MediaTable = &schema.Table{
		Name:        "media",
		Columns:     MediaColumns,
		PrimaryKey:  []*schema.Column{MediaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "media"},
	}
	// MediaFileColumns holds the columns for the "media_file" table.
	MediaFileColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "rendition_name", Type: field.TypeString, Size: 100},
		{Name: "format", Type: field.TypeString},
		{Name: "target_bandwidth", Type: field.TypeUint64, Default: 800000},
		{Name: "video_bitrate", Type: field.TypeInt64},
		{Name: "audio_bitrate", Type: field.TypeInt64},
		{Name: "video_codec", Type: field.TypeString, Default: "h264"},
		{Name: "audio_codec", Type: field.TypeString, Default: "aac"},
		{Name: "resolution_width", Type: field.TypeUint16},
		{Name: "resolution_height", Type: field.TypeUint16},
		{Name: "framerate", Type: field.TypeUint8},
		{Name: "duration_seconds", Type: field.TypeFloat64},
		{Name: "media_type", Type: field.TypeEnum, Enums: []string{"audio", "video"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"Pending", "Processing", "Ready", "Errored"}},
		{Name: "message", Type: field.TypeString, Nullable: true, Size: 255, Default: ""},
		{Name: "entry_file", Type: field.TypeString, Size: 255, Default: "index.m3u8"},
		{Name: "mimetype", Type: field.TypeString, Size: 255, Default: "application/x-mpegURL"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "media", Type: field.TypeUUID, Nullable: true},
	}
	// MediaFileTable holds the schema information for the "media_file" table.
	MediaFileTable = &schema.Table{
		Name:       "media_file",
		Columns:    MediaFileColumns,
		PrimaryKey: []*schema.Column{MediaFileColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "media_file_media_media_files",
				Columns: []*schema.Column{MediaFileColumns[19]},

				RefColumns: []*schema.Column{MediaColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "media_file"},
	}
	// MediaProbeColumns holds the columns for the "media_probe" table.
	MediaProbeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "filename", Type: field.TypeString, Size: 255},
		{Name: "filesize", Type: field.TypeInt},
		{Name: "checksum_sha256", Type: field.TypeString, Size: 128},
		{Name: "aspect_ratio", Type: field.TypeString, Size: 5, Default: "16:9"},
		{Name: "width", Type: field.TypeInt},
		{Name: "height", Type: field.TypeInt},
		{Name: "duration_seconds", Type: field.TypeFloat64},
		{Name: "video_bitrate", Type: field.TypeInt},
		{Name: "audio_bitrate", Type: field.TypeInt},
		{Name: "framerate", Type: field.TypeFloat64},
		{Name: "format", Type: field.TypeString},
		{Name: "nb_streams", Type: field.TypeInt, Default: 1},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "media", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// MediaProbeTable holds the schema information for the "media_probe" table.
	MediaProbeTable = &schema.Table{
		Name:       "media_probe",
		Columns:    MediaProbeColumns,
		PrimaryKey: []*schema.Column{MediaProbeColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "media_probe_media_probe",
				Columns: []*schema.Column{MediaProbeColumns[15]},

				RefColumns: []*schema.Column{MediaColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "media_probe"},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MediaTable,
		MediaFileTable,
		MediaProbeTable,
	}
)

func init() {
	MediaFileTable.ForeignKeys[0].RefTable = MediaTable
	MediaProbeTable.ForeignKeys[0].RefTable = MediaTable
}
