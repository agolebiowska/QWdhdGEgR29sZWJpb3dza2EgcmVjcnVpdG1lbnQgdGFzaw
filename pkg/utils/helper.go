package utils

// Helper routines that allocates a new type value
// to store v and returns a pointer to it.

func Int(v int) *int { return &v }

func Int64(v int64) *int64 { return &v }

func Float32(v float32) *float32 { return &v }

func String(v string) *string { return &v }
