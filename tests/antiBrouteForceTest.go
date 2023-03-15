package tests

import (
	"net/http"
	"test/internal"
	"test/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testAuthService(t *testing.T) {
	// Создаем тестовый сервер с методами API 
	internal.Run()

	repository.AddWhiteList("192.168.0.1")
	repository.AddBlackList("10.0.0.1")

    // Выполняем различные сценарии тестирования с помощью библиотеки testify 
	t.Run("whitelist", func(t *testing.T) {
        // Проверяем что ip из whitelist всегда получает ok=true 
        resp ,err := http.Post("localhost:3080/api")
        assert.NoError(t,err)
        assert.Equal(t,"true",resp.Body.JSON())
    })

	t.Run("blacklist", func(t *testing.T) {
        // Проверяем что ip из blacklist всегда получает ok=false 
        resp ,err := http.Post("localhost:3080/api")
        assert.NoError(t,err)
        assert.Equal(t,"false",resp.Body.JSON())
    })

	t.Run("normal", func(t *testing.T) {
         // Проверяем что обычный ip получает ok=true при нормальной частоте попыток  
         resp ,err := http.Post("localhost:3080/api")
         assert.NoError(t,err) 
		 assert.Equal(t,"true",resp.Body.JSON())
	})
	
	t.Run("bruteforce", func(t *testing.T) {
		// Проверяем что обычный ip получает ok=false при превышении частоты попыток по логину/паролю/ip 
		for i := 0; i < 11; i++ {
			resp ,err := http.Post("localhost:3080/api")
			assert.NoError(t,err)
			if i < 10 {
				assert.Equal(t, "true", resp.Body.JSON()) // первые 10 попыток разрешены
			} else {
				assert.Equal(t, "false", resp.Body.JSON()) // 11-я попытка отклонена
			}
		}
	})

}