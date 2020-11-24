/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package metadata

import (
	"fmt"
	"io/ioutil"

	"github.com/buildpacks/libcnb"
	"github.com/pelletier/go-toml"

	"github.com/buildpacks/github-actions/internal/toolkit"
)

func ComputeMetadata(tk toolkit.Toolkit) error {
	path := "buildpack.toml"
	if s, ok := tk.GetInput("path"); ok {
		path = s
	}

	c, err := ioutil.ReadFile(path)
	if err != nil {
		return toolkit.FailedErrorf("unable to read %s", path)
	}

	var bp libcnb.Buildpack
	if err := toml.Unmarshal(c, &bp); err != nil {
		return toolkit.FailedErrorf("unable to unmarshal %s", path)
	}

	fmt.Printf(`Metadata:
  ID:       %s
  Name:     %s
  Version:  %s
  Homepage: %s
`, bp.Info.ID, bp.Info.Name, bp.Info.Version, bp.Info.Homepage)

	tk.SetOutput("id", bp.Info.ID)
	tk.SetOutput("name", bp.Info.Name)
	tk.SetOutput("version", bp.Info.Version)
	tk.SetOutput("homepage", bp.Info.Homepage)

	return nil
}
