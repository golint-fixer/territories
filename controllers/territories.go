// Bundle of functions managing the CRUD and the elasticsearch engine
package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/logs"
	"github.com/quorumsco/territories/models"
)

// Territory contains the territory related methods and a gorm client
type Territory struct {
	DB *gorm.DB
}

// RetrieveCollection calls the TerritorySQL Find method and returns the results via RPC
func (t *Territory) RetrieveCollection(args models.TerritoryArgs, reply *models.TerritoryReply) error {
	logs.Debug("RetrieveCollection de Territory")
	var (
		territoryStore = models.TerritoryStore(t.DB)
		err            error
	)

	if reply.Territories, err = territoryStore.Find(args); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

// Retrieve calls the TerritorySQL First method and returns the results via RPC
func (t *Territory) Retrieve(args models.TerritoryArgs, reply *models.TerritoryReply) error {
	logs.Debug("Retrieve de Territory")
	var (
		territoryStore = models.TerritoryStore(t.DB)
		err            error
	)
	logs.Debug("Retrieve de Territory -> ça passe1")
	logs.Debug(args.Territory)
	logs.Debug(args.Territory.ID)
	logs.Debug(args.Territory.GroupID)
	if reply.Territory, err = territoryStore.First(args); err != nil {
		logs.Error("Retrieve de Territory -> ça passe pas...")
		logs.Debug(err)
		logs.Error(err)
		return err
	}

	return nil
}

// Update calls the TerritorySQL Update method and returns the results via RPC
func (t *Territory) Update(args models.TerritoryArgs, reply *models.TerritoryReply) error {
	var (
		territoryStore = models.TerritoryStore(t.DB)
		err            error
	)

	if err = territoryStore.Save(args.Territory, args); err != nil {
		logs.Error(err)
		return err
	}

	//args.Territory.Address = models.Address{}

	if reply.Territory, err = territoryStore.First(args); err != nil {
		return err
	}

	return nil
}

// Create calls the TerritorySQL Create method and returns the results via RPC
func (t *Territory) Create(args models.TerritoryArgs, reply *models.TerritoryReply) error {
	var (
		territoryStore = models.TerritoryStore(t.DB)
		err            error
	)

	if err = territoryStore.Save(args.Territory, args); err != nil {
		logs.Error(err)
		return err
	}

	reply.Territory = args.Territory

	return nil
}

// Delete calls the TerritorySQL Delete method and returns the results via RPC
func (t *Territory) Delete(args models.TerritoryArgs, reply *models.TerritoryReply) error {
	var (
		territoryStore = models.TerritoryStore(t.DB)
		err            error
	)

	if err = territoryStore.Delete(args.Territory, args); err != nil {
		logs.Debug(err)
		return err
	}

	return nil
}
