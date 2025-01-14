/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package glance

const (
	// ServiceName -
	ServiceName = "glance"
	// ServiceAccount -
	ServiceAccount = "glance-operator-glance"
	// ServiceType -
	ServiceType = "image"
	// DatabaseName -
	DatabaseName = "glance"

	// GlanceAdminPort -
	GlanceAdminPort int32 = 9292
	// GlancePublicPort -
	GlancePublicPort int32 = 9292
	// GlanceInternalPort -
	GlanceInternalPort int32 = 9292

	// KollaConfig -
	KollaConfig = "/var/lib/config-data/merged/glance-api-config.json"
)
