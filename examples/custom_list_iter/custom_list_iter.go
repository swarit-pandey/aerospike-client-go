/*
 * Copyright 2014-2021 Aerospike, Inc.
 *
 * Portions may be licensed to Aerospike, Inc. under one or more contributor
 * license agreements.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
	"log"
	"reflect"
	"time"

	as "github.com/aerospike/aerospike-client-go"
	shared "github.com/aerospike/aerospike-client-go/examples/shared"
)

var (
	ll  = [][]int64{{40937, 1388370, 1543094673, 26, 1545655173}, {40937, 523757, 1543094673, 26, 1545655173}, {40937, 639127, 1543094673, 26, 1545655173}, {16626, 516619, 1545372572, 2, 1545547652}, {16626, 516616, 1545372572, 2, 1545547652}, {8787, 489902, 1543897224, 23, 1545655164}, {8787, 489899, 1543897224, 23, 1545655164}, {8787, 408082, 1543897224, 23, 1545655164}, {8787, 407868, 1543897224, 23, 1545655164}, {20487, 963712, 1545372572, 1, 1545655172}, {20487, 856483, 1545372572, 1, 1545655172}, {20487, 690802, 1545372572, 1, 1545655172}, {40937, 773994, 1545372572, 1, 1545655172}, {40937, 773992, 1545372572, 1, 1545655172}, {40937, 773965, 1545372572, 1, 1545655172}, {40937, 773399, 1545372572, 1, 1545655172}, {40937, 773369, 1545372572, 1, 1545655172}, {40937, 671901, 1545372572, 1, 1545655172}, {40937, 368122, 1545372572, 1, 1545655172}, {40937, 368110, 1545372572, 1, 1545655172}, {40937, 368055, 1545372572, 1, 1545655172}, {40937, 368052, 1545372572, 1, 1545655172}, {40937, 368047, 1545372572, 1, 1545655172}, {40937, 368046, 1545372572, 1, 1545655172}, {40937, 367852, 1545372572, 1, 1545655172}, {40937, 367845, 1545372572, 1, 1545655172}, {40937, 367819, 1545372572, 1, 1545655172}, {40937, 367799, 1545372572, 1, 1545655172}, {40937, 366992, 1545372572, 1, 1545655172}, {40937, 358050, 1545372572, 1, 1545655172}, {40937, 356916, 1545372572, 1, 1545655172}, {40937, 356914, 1545372572, 1, 1545655172}, {40937, 356884, 1545372572, 1, 1545655172}, {40937, 24148, 1545372572, 1, 1545655172}, {40937, 14970, 1545372572, 1, 1545655172}, {27768, 1081131, 1544417593, 1, 1544417593}, {27768, 1081125, 1544417593, 1, 1544417593}, {27768, 1081117, 1544417593, 1, 1544417593}, {27768, 1081111, 1544417593, 1, 1544417593}, {27768, 614520, 1544417593, 1, 1544417593}, {27768, 550328, 1544417593, 1, 1544417593}, {16016, 516619, 1544270206, 1, 1544270206}, {16016, 516616, 1544270206, 1, 1544270206}, {11510, 516619, 1544270101, 1, 1544270101}, {11510, 516616, 1544270101, 1, 1544270101}, {11510, 287241, 1544270101, 1, 1544270101}, {11510, 179159, 1544270101, 1, 1544270101}, {11510, 158627, 1544270101, 1, 1544270101}, {11510, 158623, 1544270101, 1, 1544270101}, {11510, 142416, 1544270101, 1, 1544270101}, {9980, 1142192, 1544270101, 1, 1544270101}, {11510, 191410, 1544270101, 1, 1544270101}, {11510, 152422, 1544270101, 1, 1544270101}, {11510, 32104, 1544270101, 1, 1544270101}, {11510, 1116, 1544270101, 1, 1544270101}, {8787, 38845, 1544270101, 2, 1545655201}, {9980, 1082750, 1544270101, 2, 1545655201}, {19506, 516619, 1544270101, 2, 1545655201}, {19506, 516616, 1544270101, 2, 1545655201}, {19506, 287241, 1544270101, 2, 1545655201}, {19506, 179159, 1544270101, 2, 1545655201}, {19506, 158627, 1544270101, 2, 1545655201}, {19506, 158623, 1544270101, 2, 1545655201}, {19506, 142416, 1544270101, 2, 1545655201}, {9980, 1142179, 1544270101, 2, 1545655201}, {19506, 422943, 1544270101, 2, 1545655201}, {19506, 345569, 1544270101, 2, 1545655201}, {19506, 63617, 1544270101, 2, 1545655201}, {19506, 1116, 1544270101, 2, 1545655201}, {20932, 516619, 1543990832, 1, 1543990832}, {20932, 516616, 1543990832, 1, 1543990832}, {20932, 458514, 1543990832, 1, 1543990832}, {20932, 369986, 1543990832, 1, 1543990832}, {20932, 203333, 1543990832, 1, 1543990832}, {20932, 1116, 1543990832, 1, 1543990832}, {12057, 516619, 1543917498, 1, 1543917498}, {12057, 516616, 1543917498, 1, 1543917498}, {12057, 422964, 1543917498, 1, 1543917498}, {12057, 191957, 1543917498, 1, 1543917498}, {12057, 161051, 1543917498, 1, 1543917498}, {12057, 1116, 1543917498, 1, 1543917498}, {38311, 516619, 1543664706, 1, 1543664706}, {38311, 516616, 1543664706, 1, 1543664706}, {38311, 670098, 1543664706, 1, 1543664706}, {38311, 638284, 1543664706, 1, 1543664706}, {38311, 1116, 1543664706, 1, 1543664706}, {23179, 516619, 1543563095, 1, 1543563095}, {23179, 516616, 1543563095, 1, 1543563095}, {23179, 413003, 1543563095, 1, 1543563095}, {23179, 412965, 1543563095, 1, 1543563095}, {23179, 1116, 1543563095, 1, 1543563095}, {40937, 494200, 1543094673, 1, 1545655173}, {40937, 494199, 1543094673, 1, 1545655173}, {40937, 494194, 1543094673, 1, 1545655173}, {40937, 494152, 1543094673, 1, 1545655173}, {40937, 494147, 1543094673, 1, 1545655173}, {40937, 494135, 1543094673, 1, 1545655173}, {11673, 516619, 1541558619, 1, 1541558619}, {11673, 516616, 1541558619, 1, 1541558619}, {11673, 191573, 1541558619, 1, 1541558619}, {11673, 141981, 1541558619, 1, 1541558619}, {11673, 1116, 1541558619, 1, 1541558619}, {8787, 800649, 1540008292, 8, 1541558572}, {8787, 800648, 1540008292, 8, 1541558572}, {28570, 839888, 1541517901, 1, 1541517901}, {28570, 839887, 1541517901, 1, 1541517901}, {28570, 839877, 1541517901, 1, 1541517901}, {28570, 502177, 1541517901, 1, 1541517901}, {28570, 437522, 1541517901, 1, 1541517901}, {28570, 196361, 1541517901, 1, 1541517901}, {28570, 155579, 1541517901, 1, 1541517901}, {28570, 129720, 1541517901, 1, 1541517901}, {28570, 129716, 1541517901, 1, 1541517901}, {28570, 74510, 1541517901, 1, 1541517901}, {28570, 44089, 1541517901, 1, 1541517901}, {28570, 44078, 1541517901, 1, 1541517901}, {28570, 36454, 1541517901, 1, 1541517901}, {28570, 34133, 1541517901, 1, 1541517901}, {28570, 34127, 1541517901, 1, 1541517901}, {28570, 31089, 1541517901, 1, 1541517901}, {28570, 31038, 1541517901, 1, 1541517901}, {28570, 30789, 1541517901, 1, 1541517901}, {28570, 120446, 1541517901, 1, 1541517901}, {28570, 30740, 1541517901, 1, 1541517901}, {28570, 973402, 1540008922, 2, 1541517862}, {28570, 973401, 1540008922, 2, 1541517862}, {28570, 838084, 1540008922, 2, 1541517862}, {28570, 838080, 1540008922, 2, 1541517862}, {28570, 196353, 1540008922, 2, 1541517862}, {28570, 195519, 1540008922, 2, 1541517862}, {28570, 43878, 1540008922, 2, 1541517862}, {28570, 43876, 1540008922, 2, 1541517862}, {28570, 516619, 1540008922, 2, 1541517862}, {28570, 516616, 1540008922, 2, 1541517862}, {28570, 973303, 1540008922, 2, 1541517862}, {28570, 839958, 1540008922, 2, 1541517862}, {28570, 839864, 1540008922, 2, 1541517862}, {28570, 671901, 1540008922, 2, 1541517862}, {28570, 129719, 1540008922, 2, 1541517862}, {28570, 34115, 1540008922, 2, 1541517862}, {28570, 31026, 1540008922, 2, 1541517862}, {28570, 24148, 1540008922, 2, 1541517862}, {28570, 6463, 1540008922, 2, 1541517862}, {28570, 5915, 1540008922, 2, 1541517862}, {28569, 497873, 1540008922, 2, 1541517862}, {28569, 26031, 1540008922, 2, 1541517862}, {28569, 1116, 1540008922, 2, 1541517862}, {10071, 516619, 1540008295, 2, 1541516935}, {10071, 516616, 1540008295, 2, 1541516935}, {10071, 1122656, 1540008295, 2, 1541516935}, {10071, 188230, 1540008295, 2, 1541516935}, {10071, 26031, 1540008295, 2, 1541516935}, {10071, 1116, 1540008295, 2, 1541516935}, {67930, 516619, 1539958836, 3, 1540180956}, {67930, 516616, 1539958836, 3, 1540180956}, {67930, 1384048, 1539958836, 3, 1540180956}, {67930, 1384044, 1539958836, 3, 1540180956}, {67930, 1116, 1539958836, 3, 1540180956}, {28570, 839869, 1540008922, 1, 1540008922}, {28570, 839866, 1540008922, 1, 1540008922}, {28570, 839865, 1540008922, 1, 1540008922}, {28570, 685624, 1540008922, 1, 1540008922}, {28570, 678302, 1540008922, 1, 1540008922}, {28570, 678266, 1540008922, 1, 1540008922}, {28570, 678263, 1540008922, 1, 1540008922}, {28570, 678186, 1540008922, 1, 1540008922}, {28570, 502150, 1540008922, 1, 1540008922}, {28570, 502132, 1540008922, 1, 1540008922}, {28570, 437443, 1540008922, 1, 1540008922}, {28570, 437413, 1540008922, 1, 1540008922}, {28570, 417952, 1540008922, 1, 1540008922}, {28570, 417939, 1540008922, 1, 1540008922}, {28570, 219690, 1540008922, 1, 1540008922}, {28570, 196479, 1540008922, 1, 1540008922}, {28570, 149483, 1540008922, 1, 1540008922}, {28570, 120450, 1540008922, 1, 1540008922}, {28570, 120449, 1540008922, 1, 1540008922}, {28570, 120448, 1540008922, 1, 1540008922}, {28570, 101942, 1540008922, 1, 1540008922}, {28570, 54949, 1540008922, 1, 1540008922}, {28570, 54938, 1540008922, 1, 1540008922}, {28570, 54926, 1540008922, 1, 1540008922}, {28570, 54921, 1540008922, 1, 1540008922}, {28570, 43912, 1540008922, 1, 1540008922}, {28570, 43886, 1540008922, 1, 1540008922}, {28570, 43880, 1540008922, 1, 1540008922}, {28570, 34124, 1540008922, 1, 1540008922}, {28570, 34123, 1540008922, 1, 1540008922}, {28570, 34121, 1540008922, 1, 1540008922}, {28570, 32042, 1540008922, 1, 1540008922}, {28570, 31007, 1540008922, 1, 1540008922}, {28570, 30984, 1540008922, 1, 1540008922}, {28570, 30983, 1540008922, 1, 1540008922}, {28570, 30957, 1540008922, 1, 1540008922}, {28570, 30730, 1540008922, 1, 1540008922}, {28570, 25315, 1540008922, 1, 1540008922}, {28570, 25231, 1540008922, 1, 1540008922}, {28570, 1315277, 1540008922, 1, 1540008922}, {28570, 1315276, 1540008922, 1, 1540008922}, {28570, 1315275, 1540008922, 1, 1540008922}, {28570, 1315273, 1540008922, 1, 1540008922}, {28570, 1315272, 1540008922, 1, 1540008922}, {28570, 1084896, 1540008922, 1, 1540008922}, {28570, 1084894, 1540008922, 1, 1540008922}, {28570, 1084766, 1540008922, 1, 1540008922}, {28570, 1084765, 1540008922, 1, 1540008922}, {28570, 973557, 1540008922, 1, 1540008922}, {28570, 973497, 1540008922, 1, 1540008922}, {28570, 956784, 1540008922, 1, 1540008922}, {28570, 956775, 1540008922, 1, 1540008922}, {28570, 956774, 1540008922, 1, 1540008922}, {28570, 946292, 1540008922, 1, 1540008922}, {28570, 946286, 1540008922, 1, 1540008922}, {28570, 946267, 1540008922, 1, 1540008922}, {28570, 842371, 1540008922, 1, 1540008922}, {28570, 765557, 1540008922, 1, 1540008922}, {28570, 697212, 1540008922, 1, 1540008922}, {28570, 679835, 1540008922, 1, 1540008922}, {28570, 678299, 1540008922, 1, 1540008922}, {28570, 671925, 1540008922, 1, 1540008922}, {28570, 671924, 1540008922, 1, 1540008922}, {28570, 501723, 1540008922, 1, 1540008922}, {28570, 296271, 1540008922, 1, 1540008922}, {28570, 149481, 1540008922, 1, 1540008922}, {28570, 129721, 1540008922, 1, 1540008922}, {28570, 120439, 1540008922, 1, 1540008922}, {28570, 36903, 1540008922, 1, 1540008922}, {28570, 34128, 1540008922, 1, 1540008922}, {28570, 25320, 1540008922, 1, 1540008922}, {28570, 25253, 1540008922, 1, 1540008922}, {28570, 25222, 1540008922, 1, 1540008922}, {28570, 2090, 1540008922, 1, 1540008922}, {23594, 516619, 1540008299, 1, 1540008299}, {23594, 516616, 1540008299, 1, 1540008299}, {23594, 417719, 1540008299, 1, 1540008299}, {23594, 401813, 1540008299, 1, 1540008299}, {23594, 1116, 1540008299, 1, 1540008299}, {26358, 516619, 1540008299, 1, 1540008299}, {26358, 516616, 1540008299, 1, 1540008299}, {26358, 287241, 1540008299, 1, 1540008299}, {26358, 179159, 1540008299, 1, 1540008299}, {26358, 158627, 1540008299, 1, 1540008299}, {26358, 158623, 1540008299, 1, 1540008299}, {26358, 142416, 1540008299, 1, 1540008299}, {26358, 828103, 1540008299, 1, 1540008299}, {26358, 819629, 1540008299, 1, 1540008299}, {26358, 1113411, 1540008299, 1, 1540008299}, {26358, 453015, 1540008299, 1, 1540008299}, {26358, 453011, 1540008299, 1, 1540008299}, {26358, 1116, 1540008299, 1, 1540008299}, {9516, 516619, 1540008299, 1, 1540008299}, {9516, 516616, 1540008299, 1, 1540008299}, {9516, 422962, 1540008299, 1, 1540008299}, {9516, 187674, 1540008299, 1, 1540008299}, {9516, 3007, 1540008299, 1, 1540008299}, {9516, 1116, 1540008299, 1, 1540008299}, {16646, 516619, 1540008299, 1, 1540008299}, {16646, 516616, 1540008299, 1, 1540008299}, {16646, 279385, 1540008299, 1, 1540008299}, {16646, 244293, 1540008299, 1, 1540008299}, {16646, 1116, 1540008299, 1, 1540008299}, {43982, 516619, 1540008298, 1, 1540008298}, {43982, 516616, 1540008298, 1, 1540008298}, {43982, 833456, 1540008298, 1, 1540008298}, {43982, 832740, 1540008298, 1, 1540008298}, {43982, 1116, 1540008298, 1, 1540008298}, {7732, 516619, 1540008294, 1, 1540008294}, {7732, 516616, 1540008294, 1, 1540008294}, {8787, 1315277, 1540008294, 1, 1540008294}, {8787, 1315276, 1540008294, 1, 1540008294}, {8787, 1315275, 1540008294, 1, 1540008294}, {8787, 1315273, 1540008294, 1, 1540008294}, {8787, 1315272, 1540008294, 1, 1540008294}, {8787, 1084896, 1540008294, 1, 1540008294}, {8787, 1084894, 1540008294, 1, 1540008294}, {8787, 1084766, 1540008294, 1, 1540008294}, {8787, 1084765, 1540008294, 1, 1540008294}, {8787, 973557, 1540008294, 1, 1540008294}, {8787, 973497, 1540008294, 1, 1540008294}, {8787, 973303, 1540008294, 1, 1540008294}, {8787, 765557, 1540008294, 1, 1540008294}, {8787, 697212, 1540008294, 1, 1540008294}, {8787, 685730, 1540008294, 1, 1540008294}, {8787, 685727, 1540008294, 1, 1540008294}, {8787, 685723, 1540008294, 1, 1540008294}, {8787, 685624, 1540008294, 1, 1540008294}, {8787, 671901, 1540008294, 1, 1540008294}, {8787, 287241, 1540008294, 1, 1540008294}, {8787, 179159, 1540008294, 1, 1540008294}, {8787, 158627, 1540008294, 1, 1540008294}, {8787, 158623, 1540008294, 1, 1540008294}, {8787, 142416, 1540008294, 1, 1540008294}, {8787, 36903, 1540008294, 1, 1540008294}, {8787, 2090, 1540008294, 1, 1540008294}, {9980, 1082754, 1540008294, 1, 1540008294}, {7732, 185880, 1540008294, 1, 1540008294}, {7732, 1119, 1540008294, 1, 1540008294}, {7732, 1116, 1540008294, 1, 1540008294}, {20487, 516619, 1540008292, 1, 1540008292}, {20487, 516616, 1540008292, 1, 1540008292}, {20487, 364206, 1540008292, 1, 1540008292}, {20487, 358960, 1540008292, 1, 1540008292}, {20487, 1116, 1540008292, 1, 1540008292}, {41111, 516619, 1540008292, 3, 1540008292}, {41111, 516616, 1540008292, 3, 1540008292}, {41111, 742014, 1540008292, 3, 1540008292}, {41111, 740077, 1540008292, 3, 1540008292}, {41111, 1116, 1540008292, 3, 1540008292}, {29860, 516619, 1538708906, 1, 1538708906}, {29860, 516616, 1538708906, 1, 1538708906}, {29860, 514572, 1538708906, 1, 1538708906}, {29860, 496357, 1538708906, 1, 1538708906}, {29860, 1116, 1538708906, 1, 1538708906}}
	cll = SliceListInt64(ll)
)

type SliceInt64 []int64

func (sli SliceInt64) Len() int {
	return len(sli)
}

func (sli SliceInt64) PackList(buf as.BufferEx) (int, error) {
	size := 0
	for _, elem := range sli {
		n, err := as.PackInt64(buf, elem)
		size += n
		if err != nil {
			return size, err
		}
	}
	return size, nil
}

type SliceListInt64 [][]int64

func (sli SliceListInt64) Len() int {
	return len(sli)
}

func (sli SliceListInt64) PackList(buf as.BufferEx) (int, error) {
	size := 0
	for _, l := range sli {
		n, err := as.PackList(buf, SliceInt64(l))
		size += n
		if err != nil {
			return size, err
		}
	}
	return size, nil
}

func main() {
	runExample(shared.Client)
	log.Println("Example finished successfully.")
}

func runExample(client *as.Client) {
	key, err := as.NewKey(*shared.Namespace, *shared.Set, "addkey")
	shared.PanicOnError(err)

	// Delete record if it already exists.
	client.Delete(shared.WritePolicy, key)

	bin := as.NewBin("bin", ll)
	b := time.Now()
	for i := 0; i < 10000; i++ {
		if err := client.PutBins(shared.WritePolicy, key, bin); err != nil {
			log.Panicf("Get failed %s: %s", key, err)
		}
	}
	t := time.Since(b)
	log.Println("Performance WITHOUT custom iterator:", t)

	bin = as.NewBin("bin", cll)
	b = time.Now()
	for i := 0; i < 10000; i++ {
		if err := client.PutBins(shared.WritePolicy, key, bin); err != nil {
			log.Panicf("Get failed %s: %s", key, err)
		}
	}
	t = time.Since(b)
	log.Println("Performance WITH    custom iterator:", t)

	record, err := client.Get(shared.Policy, key, bin.Name)
	shared.PanicOnError(err)

	if record == nil {
		log.Fatalf(
			"Failed to get: namespace=%s set=%s key=%s",
			key.Namespace(), key.SetName(), key.Value())
	}

	// The value received from the server is an unsigned byte stream.
	// Convert to an integer before comparing with expected.
	received := [][]int64{} // returns as [][]interface{}, need to be converted
	for _, l1 := range record.Bins[bin.Name].([]interface{}) {
		l := []int64{}
		for _, l2 := range l1.([]interface{}) {
			l = append(l, int64(l2.(int)))
		}
		received = append(received, l)
	}
	expected := ll

	if reflect.DeepEqual(received, expected) {
		log.Printf("Packing successful: ns=%s set=%s key=%s bin=%s",
			key.Namespace(), key.SetName(), key.Value(), bin.Name)
	} else {
		log.Fatalf("Packing mismatch: Expected %#v.\n Received %#v.", expected, received)
	}
}
