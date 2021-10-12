package service

// import (
// 	"os"
// 	"path/filepath"
// 	"time"
// )

// func (g *GravacaoService) beginHousekeeper() {
// 	d := time.Now().Add(time.Duration(-g.horasRetencao) * time.Hour)

// 	err := filepath.Walk(g.armazenamento, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err // TODO testar com diretorio errado...
// 		}

// 		if info.ModTime().Before(d) {
// 			os.Remove(path)
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		g.errChan <- err
// 	}
// }
