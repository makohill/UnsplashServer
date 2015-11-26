package main

import "fmt"

var currentId int

var unsplashs Unsplashs

// Give us some seed data
func init() {
	RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
	RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
	RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
	RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
	RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
	RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
	RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})
    RepoCreateUnsplash(Unsplash{Name: "Gian-Reto Tarnutzer",FullUrl:"https://images.unsplash.com/photo-1447522200268-a0378dac3fba?crop=entropy&fit=crop&fm=jpg&h=650&ixjsv=2.0.0&ixlib=rb-0.3.5&q=80"})

}

func RepoFindUnsplash(id int) Unsplash {
	for _, t := range unsplashs {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Unsplash{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateUnsplash(t Unsplash) Unsplash {
	currentId += 1
	t.Id = currentId
	unsplashs = append(unsplashs, t)
	return t
}

func RepoDestroyUnsplash(id int) error {
	for i, t := range unsplashs {
		if t.Id == id {
			unsplashs = append(unsplashs[:i], unsplashs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
