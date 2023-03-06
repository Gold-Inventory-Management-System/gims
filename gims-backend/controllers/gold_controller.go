package controllers

import (
	"fmt"
	"net/http"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/gin-gonic/gin"
)

type goldController struct {
	goldUseCase  domains.GoldUseCase
}

func NewGoldController(gu domains.GoldUseCase) *goldController {
	return &goldController{
		goldUseCase:  gu,
	}
}

func (gc *goldController) NewGold(c *gin.Context) {
	newGold := new(models.InputNewGoldDetail)
	if err := c.BindJSON(newGold); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.NewGold(newGold); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) AddGold(c *gin.Context) {
	newGoldInventory := new(models.InputNewGoldInventory)
	if err := c.BindJSON(newGoldInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.AddGold(newGoldInventory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) GetAllGoldDetail(c *gin.Context) {
	goldDetails, err := gc.goldUseCase.GetAllGoldDetail()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": goldDetails,
	})
}

func (gc *goldController) FindGoldDetailByCode(c *gin.Context) {
	code := c.Param("code")
	res, err := gc.goldUseCase.FindGoldDetailByCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (gc *goldController) FindGoldDetailByDetail(c *gin.Context) {
	goldDetail := new(models.GoldDetail)
	if err := c.BindJSON(goldDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	res, err := gc.goldUseCase.FindGoldDetailByDetail(goldDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (gc *goldController) EditGoldDetail(c *gin.Context) {
	goldDetail := new(models.GoldDetail)
	if err := c.BindJSON(goldDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.EditGoldDetail(goldDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) GetAllGoldDetailJoinInventory(c *gin.Context) {
	res, err := gc.goldUseCase.GetAllGoldDetailJoinInventory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (gc *goldController) SetStatusGoldDetailToDelete(c *gin.Context) {
	body := new(models.InputSetStatusGoldDetail)
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.SetStatusGoldDetailToDelete(body.GoldDetailID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) SetStatusGoldDetailToNormal(c *gin.Context) {
	idString := c.Param("id")
	idUint32, err := gc.goldUseCase.ConvertIDStringToUint32(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.SetStatusGoldDetailToNormal(idUint32); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) SetStatusGoldInventory(c *gin.Context) {
	input := new(models.InputSetStatusGoldInventory)
	if err := c.BindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if input.Status != "front" && input.Status != "safe" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "status must be front or back",
		})
		return
	}
	if err := gc.goldUseCase.SetStatusGoldInventory(input.GoldInventoryID, input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) GetGoldDetailJoinInventoryByDetail(c *gin.Context) {
	goldDetail := new(models.GoldDetail)
	if err := c.BindJSON(goldDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	res, err := gc.goldUseCase.GetGoldDetailJoinInventoryByDetail(goldDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": res,
	})
}

func (gc *goldController) GetGoldDetailByGoldDetailID(c *gin.Context) {
	id := c.Param("id")
	goldDetailID, err := gc.goldUseCase.ConvertIDStringToUint32(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	goldDetail, err := gc.goldUseCase.GetGoldDetailByGoldDetailID(goldDetailID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": goldDetail,
	})
}

func (gc *goldController) DeleteGoldInventoryByIDArray(c *gin.Context) {
	input := new(models.InputDeleteGoldInventory)
	if err := c.BindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.DeleteGoldInventoryByIDArray(input.GoldInventoryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) SetTagSerialNumberGoldInventory(c *gin.Context) {
	input := new(models.InputSetTagSerialNumber)
	if err := c.BindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.SetTagSerialNumberGoldInventory(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) GetGoldByTagSerailNumber(c *gin.Context) {
	sn := c.Query("serial-number")
	fmt.Println(sn)
	sn64, err := gc.goldUseCase.ConvertIDStringToUint64(sn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if sn64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "serial number must not equal to 0",
		})
		return
	}
	gold, err := gc.goldUseCase.QueryGoldByTagSerialNumber(sn64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": gold,
	})
}

func (gc *goldController) CheckGold(c *gin.Context) {
	tsa := new(models.CheckGoldBody)
	if err := c.BindJSON(tsa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"message": err.Error(),
		})
		return
	}
	res, err := gc.goldUseCase.CheckFrontGold(tsa.TagSerialArray)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "ok",
		"result": res,
	})
}

func (gc *goldController) GetAllFrontGold(c *gin.Context) {
	goldJoins, err := gc.goldUseCase.GetAllFrontGold()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "ok",
		"data": goldJoins,
	})
}