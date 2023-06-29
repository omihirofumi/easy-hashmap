# About
チェイン法を使用した、HashMapの簡単な実装
※ジェネリクスを使用しているため、Go 1.18以上で使用可能

# Installing
### Using *go get*

    $ go get github.com/omihirofumi/easy-hashmap

## Example
    import (
    	"github.com/omihirofumi/easy-hashmap/hashmap"
    )
    
    func main() {
    	m := hashmap.NewHashMap[int]()
    	m.Put("key1", 1)
    	m.Put("key2", 2)
    
    	if v1, ok := m.Get("key1"); ok {
    		...
    	}
    }
