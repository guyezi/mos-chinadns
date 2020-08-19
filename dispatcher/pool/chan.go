//     Copyright (C) 2020, IrineSistiana
//
//     This file is part of mos-chinadns.
//
//     mos-chinadns is free software: you can redistribute it and/or modify
//     it under the terms of the GNU General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     mos-chinadns is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU General Public License for more details.
//
//     You should have received a copy of the GNU General Public License
//     along with this program.  If not, see <https://www.gnu.org/licenses/>.

package pool

import (
	"github.com/miekg/dns"
	"sync"
)

var resChanPool = sync.Pool{
	New: func() interface{} {
		return make(chan *dns.Msg, 1)
	},
}

func GetResChan() chan *dns.Msg {
	return resChanPool.Get().(chan *dns.Msg)
}

func ReleaseResChan(c chan *dns.Msg) {
	for {
		select {
		case <-c:
		default:
			resChanPool.Put(c)
			return
		}
	}
}
