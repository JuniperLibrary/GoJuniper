package basics_test

import (
	"testing"

	"gojuniper/internal/basics"
)

func TestWeekdayEnum(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{"Sunday", basics.Sunday, 0},
		{"Monday", basics.Monday, 1},
		{"Tuesday", basics.Tuesday, 2},
		{"Wednesday", basics.Wednesday, 3},
		{"Thursday", basics.Thursday, 4},
		{"Friday", basics.Friday, 5},
		{"Saturday", basics.Saturday, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("%s = %d, want %d", tt.name, tt.value, tt.want)
			}
		})
	}
}

func TestStatusEnum(t *testing.T) {
	// 测试跳过0的枚举（从1开始）
	if basics.StatusPending != 1 {
		t.Errorf("StatusPending = %d, want 1", basics.StatusPending)
	}
	if basics.StatusRunning != 2 {
		t.Errorf("StatusRunning = %d, want 2", basics.StatusRunning)
	}
	if basics.StatusCompleted != 3 {
		t.Errorf("StatusCompleted = %d, want 3", basics.StatusCompleted)
	}
	if basics.StatusFailed != 4 {
		t.Errorf("StatusFailed = %d, want 4", basics.StatusFailed)
	}
}

func TestGetStatusName(t *testing.T) {
	tests := []struct {
		status int
		want   string
	}{
		{basics.StatusPending, "Pending"},
		{basics.StatusRunning, "Running"},
		{basics.StatusCompleted, "Completed"},
		{basics.StatusFailed, "Failed"},
		{999, "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := basics.GetStatusName(tt.status); got != tt.want {
				t.Errorf("GetStatusName(%d) = %s, want %s", tt.status, got, tt.want)
			}
		})
	}
}

func TestPermissionFlags(t *testing.T) {
	// 测试位运算枚举值
	if basics.PermissionRead != 1 {
		t.Errorf("PermissionRead = %d, want 1", basics.PermissionRead)
	}
	if basics.PermissionWrite != 2 {
		t.Errorf("PermissionWrite = %d, want 2", basics.PermissionWrite)
	}
	if basics.PermissionExecute != 4 {
		t.Errorf("PermissionExecute = %d, want 4", basics.PermissionExecute)
	}
}

func TestCheckPermission(t *testing.T) {
	// 组合权限：读+写
	perms := basics.PermissionRead | basics.PermissionWrite

	if !basics.CheckPermission(perms, basics.PermissionRead) {
		t.Error("expected Read permission to be set")
	}
	if !basics.CheckPermission(perms, basics.PermissionWrite) {
		t.Error("expected Write permission to be set")
	}
	if basics.CheckPermission(perms, basics.PermissionExecute) {
		t.Error("expected Execute permission NOT to be set")
	}
}

func TestCombinePermissions(t *testing.T) {
	// 测试权限组合
	combined := basics.CombinePermissions(
		basics.PermissionRead,
		basics.PermissionWrite,
	)

	// 读+写 = 1 + 2 = 3
	if combined != 3 {
		t.Errorf("combined = %d, want 3", combined)
	}

	// 三种权限组合
	all := basics.CombinePermissions(
		basics.PermissionRead,
		basics.PermissionWrite,
		basics.PermissionExecute,
	)
	if all != 7 {
		t.Errorf("all permissions = %d, want 7", all)
	}
}

func TestDirectionEnum(t *testing.T) {
	tests := []struct {
		direction basics.Direction
		value     int
		str       string
	}{
		{basics.North, 0, "North"},
		{basics.South, 1, "South"},
		{basics.East, 2, "East"},
		{basics.West, 3, "West"},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if int(tt.direction) != tt.value {
				t.Errorf("%s value = %d, want %d", tt.str, tt.direction, tt.value)
			}
			if tt.direction.String() != tt.str {
				t.Errorf("String() = %s, want %s", tt.direction.String(), tt.str)
			}
		})
	}
}

func TestColorEnum(t *testing.T) {
	tests := []struct {
		color    basics.Color
		value    int
		str      string
		r, g, b  uint8
	}{
		{basics.ColorNone, 0, "None", 0, 0, 0},
		{basics.ColorRed, 1, "Red", 255, 0, 0},
		{basics.ColorGreen, 2, "Green", 0, 255, 0},
		{basics.ColorBlue, 3, "Blue", 0, 0, 255},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if int(tt.color) != tt.value {
				t.Errorf("%s value = %d, want %d", tt.str, tt.color, tt.value)
			}
			if tt.color.String() != tt.str {
				t.Errorf("String() = %s, want %s", tt.color.String(), tt.str)
			}
			r, g, b, a := tt.color.RGBA()
			if r != tt.r || g != tt.g || b != tt.b {
				t.Errorf("RGBA() = (%d,%d,%d), want (%d,%d,%d)", r, g, b, tt.r, tt.g, tt.b)
			}
			if a != 255 {
				t.Errorf("Alpha = %d, want 255", a)
			}
		})
	}
}

func TestSeasonEnum(t *testing.T) {
	tests := []struct {
		season basics.Season
		value  int
		str    string
	}{
		{basics.SeasonSpring, 1, "Spring"},
		{basics.SeasonSummer, 2, "Summer"},
		{basics.SeasonAutumn, 3, "Autumn"},
		{basics.SeasonWinter, 4, "Winter"},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if int(tt.season) != tt.value {
				t.Errorf("%s value = %d, want %d", tt.str, tt.season, tt.value)
			}
			if tt.season.String() != tt.str {
				t.Errorf("String() = %s, want %s", tt.season.String(), tt.str)
			}
		})
	}
}

func TestFileSizeEnum(t *testing.T) {
	// 测试文件大小枚举值
	if basics.KB != 1024 {
		t.Errorf("KB = %d, want 1024", basics.KB)
	}
	if basics.MB != 1048576 {
		t.Errorf("MB = %d, want 1048576", basics.MB)
	}
	if basics.GB != 1073741824 {
		t.Errorf("GB = %d, want 1073741824", basics.GB)
	}
}

func TestFormatFileSize(t *testing.T) {
	tests := []struct {
		bytes int64
		want  string
	}{
		{500, "500 B"},
		{1024, "1.00 KB"},
		{1536, "1.50 KB"},
		{1048576, "1.00 MB"},
		{1572864, "1.50 MB"},
		{1073741824, "1.00 GB"},
		{1099511627776, "1.00 TB"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := basics.FormatFileSize(tt.bytes)
			if got != tt.want {
				t.Errorf("FormatFileSize(%d) = %s, want %s", tt.bytes, got, tt.want)
			}
		})
	}
}

func TestGetWeekdayValue(t *testing.T) {
	tests := []struct {
		day  int
		want int
	}{
		{0, basics.Sunday},
		{1, basics.Monday},
		{2, basics.Tuesday},
		{3, basics.Wednesday},
		{4, basics.Thursday},
		{5, basics.Friday},
		{6, basics.Saturday},
		{-1, -1}, // 无效值
		{7, -1},  // 无效值
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := basics.GetWeekdayValue(tt.day); got != tt.want {
				t.Errorf("GetWeekdayValue(%d) = %d, want %d", tt.day, got, tt.want)
			}
		})
	}
}

func TestFileModeEnum(t *testing.T) {
	// 测试文件模式位运算枚举
	if basics.ModeUserRead != 1 {
		t.Errorf("ModeUserRead = %d, want 1", basics.ModeUserRead)
	}
	if basics.ModeUserWrite != 8 {
		t.Errorf("ModeUserWrite = %d, want 8", basics.ModeUserWrite)
	}
	if basics.ModeUserExec != 64 {
		t.Errorf("ModeUserExec = %d, want 64", basics.ModeUserExec)
	}
}