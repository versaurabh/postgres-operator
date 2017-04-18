package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/motomux/pretty"
	"k8s.io/client-go/pkg/api/v1"

	"github.bus.zalan.do/acid/postgres-operator/pkg/spec"
	"github.bus.zalan.do/acid/postgres-operator/pkg/util/constants"
)

var passwordChars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(int64(time.Now().Unix()))
}

func RandomPassword(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = passwordChars[rand.Intn(len(passwordChars))]
	}

	return string(b)
}

func NameFromMeta(meta v1.ObjectMeta) spec.NamespacedName {
	return spec.NamespacedName{
		Namespace: meta.Namespace,
		Name:      meta.Name,
	}
}

func PodClusterName(pod *v1.Pod) spec.NamespacedName {
	if name, ok := pod.Labels[constants.ClusterNameLabel]; ok {
		return spec.NamespacedName{
			Namespace: pod.Namespace,
			Name:      name,
		}
	}

	return spec.NamespacedName{}
}

func PodSpiloRole(pod *v1.Pod) string {
	return pod.Labels[constants.SpiloRoleLabel]
}

func PGUserPassword(user spec.PgUser) string {
	s := md5.Sum([]byte(user.Password + user.Name))

	return "md5" + hex.EncodeToString(s[:])
}

func Pretty(x interface{}) (f fmt.Formatter) {
	return pretty.Formatter(x)
}

func PrettyDiff(a, b interface{}) (result string) {
	diff := pretty.Diff(a, b)
	return strings.Join(diff, "\n")
}
