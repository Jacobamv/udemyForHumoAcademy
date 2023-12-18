package mock

import (
	"udemy/pkg"
	"udemy/pkg/models"
)

func FillRegions() {
	pkg.Regions = append(pkg.Regions, models.Region{
		Name: "Хатлон",
	})

	pkg.Regions = append(pkg.Regions, models.Region{
		Name: "Согд",
	})

	pkg.Regions = append(pkg.Regions, models.Region{
		Name: "РРП",
	})

	pkg.Regions = append(pkg.Regions, models.Region{
		Name: "ГБАО",
	})
}
