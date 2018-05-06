package main

import (
  "fmt"
  "log"
  "strconv"

  "github.com/dgraph-io/badger"
)

type Database struct {
  opts   badger.Options
  name   string
}

func New(name string) *Database {
  return &Database{ name: name }
}

func (d *Database) Init() {
  d.opts = badger.DefaultOptions
  d.opts.Dir = "badger/" + d.name
  d.opts.ValueDir = "badger/" + d.name
}

func (d *Database) FindOne(key string) (string, error) {
  hit := ""
  db, err := badger.Open(d.opts)
  if err != nil {
    log.Println(d.name, "open error: ", err.Error())
    return "", err
  }
  defer db.Close()
  err = db.View(func(txn *badger.Txn) error {
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
    log.Println(d.name, "View error: ", err.Error())
    return "", err
  }

  fmt.Println(key, hit)
  return hit, nil
}

func (d *Database) FindAll() []Nero {
  res := []Nero{}

  db, err := badger.Open(d.opts)

  if err != nil {
    log.Println(d.name, "open error: ", err.Error())
    return res
  }

  defer db.Close()

  db.View(func(txn *badger.Txn) error {
    opts := badger.DefaultIteratorOptions
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

func (d *Database) Update(key string, value string) error {
  db, err := badger.Open(d.opts)

  if err != nil {
    log.Println(d.name, "open error: ", err.Error())
    return err
  }

  defer db.Close()

  return db.Update(func(tx *badger.Txn) error {
    return tx.Set([]byte(key), []byte(value))
  })
}

func (d *Database) UpdateTxn(key string, value string) error {
  db, err := badger.Open(d.opts)

  if err != nil {
    log.Println(d.name, "open error: ", err.Error())
    return err
  }

  defer db.Close()

  return db.Update(func(tx *badger.Txn) error {
    txn := db.NewTransaction(true)

    defer txn.Discard()

    if err := txn.Set([]byte(key), []byte(value)); err != nil {
      return err
    }

    return txn.Commit(nil)

  })
}

func (d *Database) ResetAll(value string) error {
  db, err := badger.Open(d.opts)

  if err != nil {
    log.Println(d.name, "open error: ", err.Error())
    return err
  }

  defer db.Close()

  keys := []string{}

  db.View(func(txn *badger.Txn) error {
    opts := badger.DefaultIteratorOptions
    opts.PrefetchValues = false
    it := txn.NewIterator(opts)
    defer it.Close()
    for it.Rewind(); it.Valid(); it.Next() {
      item := it.Item()
      k := item.Key()
      keys = append(keys, string(k))
    }
    return nil
  })

  return db.Update(func(tx *badger.Txn) error {
    for _, v := range keys {
      tx.Set([]byte(v), []byte(value))
    }
    return nil
  })
}
