// Bundle of functions managing the CRUD and the elasticsearch engine
package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/logs"
	"github.com/quorumsco/territories/models"
)

// Note contains the note related methods and a gorm client
type PointLocation struct {
	DB *gorm.DB
}

// RetrieveCollection calls the NoteSQL Find method and returns the results via RPC
func (t *PointLocation) RetrieveCollection(args models.PointLocationArgs, reply *models.PointLocationReply) error {
	logs.Debug("RetrieveCollection de POINTLOCATION->ça passe pour l'instant...1")
	var (
		err error

		PointLocationStore = models.PointLocationStore(t.DB)
	)
	logs.Debug("RetrieveCollection de POINTLOCATION->ça passe pour l'instant...2")
	reply.Polygon, err = PointLocationStore.Find(args)
	if err != nil {
		logs.Debug("ouchhhh")
		logs.Error(err)
		return err
	}

	return nil
}

// Retrieve calls the NoteSQL First method and returns the results via RPC
func (t *PointLocation) Retrieve(args models.PointLocationArgs, reply *models.PointLocationReply) error {
	var (
		PointLocationStore = models.PointLocationStore(t.DB)
		err                error
	)

	if reply.PointLocation, err = PointLocationStore.First(args); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

// Create calls the NoteSQL Save method and returns the results via RPC
func (t *PointLocation) Create(args models.PointLocationArgs, reply *models.PointLocationReply) error {
	var (
		err error

		PointLocationStore = models.PointLocationStore(t.DB)
	)

	if err = PointLocationStore.Save(args.PointLocation, args); err != nil {
		logs.Error(err)
		return err
	}

	reply.PointLocation = args.PointLocation

	return nil
}

// Delete calls the NoteSQL Delete method and returns the results via RPC
func (t *PointLocation) Delete(args models.PointLocationArgs, reply *models.PointLocationReply) error {
	var (
		err error

		PointLocationStore = models.PointLocationStore(t.DB)
	)

	if err = PointLocationStore.Delete(args.PointLocation, args); err != nil {
		logs.Debug(err)
		return err
	}

	return nil
}
