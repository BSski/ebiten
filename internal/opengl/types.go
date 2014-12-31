// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opengl

type FilterType int
type ShaderType int
type BufferType int
type BufferUsageType int

type Context struct {
	Nearest            FilterType
	Linear             FilterType
	VertexShader       ShaderType
	FragmentShader     ShaderType
	ArrayBuffer        BufferType
	ElementArrayBuffer BufferType
	DynamicDraw        BufferUsageType
	StaticDraw         BufferUsageType
	*context
}
