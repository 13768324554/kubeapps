package cmd

import (
	"path/filepath"

	"github.com/ksonnet/kubecfg/pkg/kubecfg"
	"github.com/spf13/cobra"
	"k8s.io/client-go/pkg/api"
)

var downCmd = &cobra.Command{
	Use:   "down FLAG",
	Short: "uninstall KubeApps components",
	Long:  `uninstall KubeApps components`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := kubecfg.DeleteCmd{}
		ns, err := cmd.Flags().GetString("namespace")
		if err != nil {
			return err
		}
		c.DefaultNamespace = ns

		c.GracePeriod, err = cmd.Flags().GetInt64("grace-period")
		if err != nil {
			return err
		}

		c.ClientPool, c.Discovery, err = restClientPool()
		if err != nil {
			return err
		}

		kaManifestDir, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}
		if kaManifestDir == "" {
			home, err := getHome()
			if err != nil {
				return err
			}
			kaManifestDir = filepath.Join(home, KUBEAPPS_DIR)
		}

		objs, err := parseObjects(kaManifestDir)
		if err != nil {
			return err
		}
		return c.Run(objs)
	},
}

func init() {
	RootCmd.AddCommand(downCmd)
	downCmd.Flags().StringP("namespace", "", api.NamespaceDefault, "Specify namespace for the KubeApps components")
	downCmd.Flags().Int64("grace-period", -1, "Number of seconds given to resources to terminate gracefully. A negative value is ignored")
	downCmd.Flags().String("path", "", "Specify folder contains the manifests")
}
