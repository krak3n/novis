package novis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNovis_Rev(t *testing.T) {
	type tc struct {
		tname  string
		novis  *Novis
		name   string
		values []string
		path   string
	}
	tt := []tc{
		{
			tname: "nil root",
			novis: &Novis{},
			path:  "",
		},
		{
			tname: "empty root",
			novis: New(),
			path:  "",
		},
		{
			tname: "empty name",
			novis: &Novis{
				Root: &Branch{
					branches: map[string]*Branch{
						"foo": &Branch{
							name: "foo",
							path: "/foo",
						},
					},
				},
			},
			path: "",
		},
		{
			tname: "reverse without params",
			novis: &Novis{
				Root: &Branch{
					branches: map[string]*Branch{
						"foo": &Branch{
							name: "foo",
							path: "/foo",
							branches: map[string]*Branch{
								"bar": &Branch{
									name: "bar",
									path: "/bar",
								},
							},
						},
					},
				},
			},
			name: "foo.bar",
			path: "/foo/bar",
		},
		{
			tname: "reverse with params",
			novis: &Novis{
				Root: &Branch{
					branches: map[string]*Branch{
						"foo": &Branch{
							name:   "foo",
							path:   "/foo/:id",
							params: []string{":id"},
							branches: map[string]*Branch{
								"bar": &Branch{
									name: "bar",
									path: "/bar",
								},
							},
						},
					},
				},
			},
			name:   "foo.bar",
			values: []string{"9fd4f91a-115d-4be6-a3d1-00925f75db1f"},
			path:   "/foo/9fd4f91a-115d-4be6-a3d1-00925f75db1f/bar",
		},
		{
			tname: "reverse too many params",
			novis: &Novis{
				Root: &Branch{
					branches: map[string]*Branch{
						"foo": &Branch{
							name:   "foo",
							path:   "/foo/:id",
							params: []string{":id"},
							branches: map[string]*Branch{
								"bar": &Branch{
									name: "bar",
									path: "/bar",
								},
							},
						},
					},
				},
			},
			name:   "foo.bar",
			values: []string{"9fd4f91a-115d-4be6-a3d1-00925f75db1f", "foo", "bar"},
			path:   "/foo/9fd4f91a-115d-4be6-a3d1-00925f75db1f/bar",
		},
		{
			tname: "duplicate params",
			novis: &Novis{
				Root: &Branch{
					branches: map[string]*Branch{
						"foo": &Branch{
							name:   "foo",
							path:   "/foo/:id",
							params: []string{":id"},
							branches: map[string]*Branch{
								"bar": &Branch{
									name: "bar",
									path: "/bar/:id",
								},
							},
						},
					},
				},
			},
			name:   "foo.bar",
			values: []string{"9fd4f91a-115d-4be6-a3d1-00925f75db1f"},
			path:   "/foo/9fd4f91a-115d-4be6-a3d1-00925f75db1f/bar/9fd4f91a-115d-4be6-a3d1-00925f75db1f",
		},
		{
			tname: "nested params",
			novis: &Novis{
				Root: &Branch{
					branches: map[string]*Branch{
						"foo": &Branch{
							name:   "foo",
							path:   "/foo/:p1",
							params: []string{":p1"},
							branches: map[string]*Branch{
								"bar": &Branch{
									name:   "bar",
									path:   "/bar/:p2",
									params: []string{":p2"},
									branches: map[string]*Branch{
										"baz": &Branch{
											name:   "baz",
											path:   "/baz/:p3",
											params: []string{":p3"},
										},
									},
								},
							},
						},
					},
				},
			},
			name:   "foo.bar.baz",
			values: []string{"1", "2", "3"},
			path:   "/foo/1/bar/2/baz/3",
		},
	}
	for _, tc := range tt {
		t.Run(tc.tname, func(t *testing.T) {
			path := tc.novis.Rev(tc.name, tc.values...)
			assert.Equal(t, tc.path, path)
		})
	}
}
