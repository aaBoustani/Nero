package main

import (
  "log"

  "github.com/dgraph-io/badger"
  "fmt"
  "strconv"
)

var db *badger.DB

func Init() *badger.DB {
  opts := badger.DefaultOptions
  opts.Dir = "badger"
  opts.ValueDir = "badger"
  db, err := badger.Open(opts)
  if err != nil {
    log.Fatal(err)
  }

  return db
}

func Find(key string) (string, error) {
  hit := ""
  db := Init()
  defer db.Close()
  err := db.View(func(txn *badger.Txn) error {
    item, err := txn.Get([]byte(key))
    if err != nil {
      if err == badger.ErrKeyNotFound {
        return nil
      }
      return err
    }
    val, err := item.Value()
    if err != nil {
      return err
    }
    hit = string(val)
    return nil
  })

  if err != nil {
    log.Panic(err)
    return "", err
  }

  fmt.Println(key)
  fmt.Println(hit)
  return hit, nil
}

func Add(key string, value string) error {
  db := Init()
  defer db.Close()
  err := db.Update(func(tx *badger.Txn) error {
    // Start a writable transaction.
    txn := db.NewTransaction(true)

    defer txn.Discard()

    // Use the transaction...
    fmt.Printf("value to add is %s\n", value)
    err := txn.Set([]byte(key), []byte(value))
    if err != nil {
      return err
    }

    // Commit the transaction and check for error.
    if err := txn.Commit(nil); err != nil {
      return err
    }
    return nil
  })

  return err
}

func PrintAll() []Nero {
  db := Init()
  defer db.Close()
  res := []Nero{}
  db.View(func(txn *badger.Txn) error {
    opts := badger.DefaultIteratorOptions
    opts.PrefetchSize = 10
    it := txn.NewIterator(opts)
    defer it.Close()
    for it.Rewind(); it.Valid(); it.Next() {
      item := it.Item()
      k := item.Key()
      v, err := item.Value()
      if err != nil {
        return err
      }
      user := string(k)
      amount, err := strconv.Atoi(string(v))
      res = append(res, Nero{ User: user, Amount: amount })
      fmt.Printf("key=%s, value=%s\n", k, v)
    }
    return nil
  })
  return res
}