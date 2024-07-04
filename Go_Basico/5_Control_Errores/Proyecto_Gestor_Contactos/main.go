package main

import (
    
      "fmt"
      "bufio"
      "encoding/json"
      "os"
      "github.com/fatih/color"
)

// Estructura de contatos
type Contact struct {
  Name string `json:"name"`
  Email string `json:"email"`
  Phone string `json:"phone"`
}

// Guardar contactos en un archivo json
func saveContactsFile(contacts []Contact) error {
  file, err := os.Create("contacts.json")
  if err != nil {
    return err
  }
  
  defer file.Close()
  encoder := json.NewEncoder(file)
  err = encoder.Encode(contacts)
  if err != nil {
    return err
  }
  
  return nil

}

// Carga contactos desde un archivo json
func loadContactsFromFile(contacts *[]Contact) error {
  file, err := os.Open("contacts.json")
  if err != nil {
    return err
  }  

  defer file.Close() // Cierra el archivo abierto

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&contacts)
  if err != nil {
    return err
  }
  
  return nil

}

// Eliminar contacto por índice
func deleteContact(contacts *[]Contact, index int) error {
    if index < 0 || index >= len(*contacts) {
        return fmt.Errorf("índice fuera de rango")
    }
    *contacts = append((*contacts)[:index], (*contacts)[index+1:]...)
    return nil
}



func main() {
    // Slice de contactos
    var contacts []Contact

    // Cargar contactos existentes desde el archivo
    err := loadContactsFromFile(&contacts)
    if err != nil {
      fmt.Println("Error al cargar los contactos:", err)
    }

    // Crear instancia de fubio
    reader := bufio.NewReader(os.Stdin)

      // Crear funciones de color
      red := color.New(color.FgRed).SprintFunc()
      green := color.New(color.FgGreen).SprintFunc()
      blue := color.New(color.FgBlue).SprintFunc()
      yellow := color.New(color.FgYellow).SprintFunc()
      cyan := color.New(color.FgCyan).SprintFunc()
      magenta := color.New(color.FgMagenta).SprintFunc()
      white := color.New(color.FgWhite).SprintFunc()

    for {
        // Mostrar opciones al usuario
        fmt.Println(" ")
        fmt.Print(green(" ===== GESTOR DE CONTACTOS =====    \n",
        "1. Agregar un contacto\n",
        "2. Mostrar todos los contactos\n",
        "3. Eliminar un contacto\n",
        "4. Salir\n",
        "Elige una opción: "))

        // Leer la opcion del usuario  
        var option int 
        _, err = fmt.Scanln(&option)
        if err != nil {
          fmt.Println(red("Error al lerr la opción:", err))
          return
        }

        // Manejar la opción del usuario
        switch option {
        case 1:
            // Ingresar y crear contacto
            var c Contact
            fmt.Print(blue("Nombre: "))
            c.Name, _ = reader.ReadString('\n')
            fmt.Print(blue("Email: "))
            c.Email, _ = reader.ReadString('\n')
            fmt.Print(blue("Teléfono: "))
            c.Phone, _ = reader.ReadString('\n')

            // Limpiar caracteres de nueva línea
            c.Name = c.Name[:len(c.Name)-1]
            c.Email = c.Email[:len(c.Email)-1]
            c.Phone = c.Phone[:len(c.Phone)-1]
            
            // Agregar un contacto a Slice
            contacts = append(contacts, c)

            // Guardar en un archivo json
            if err := saveContactsFile(contacts); err != nil {
              fmt.Println(magenta("Error al guardar el contacto: ", err))
            }

        case 2:
            // Mostrar todos los Contactos
            fmt.Println(yellow("==============================================================="))
            for index, contact := range contacts {
              fmt.Printf("%d. Nombre: %s Email: %s Teléfono: %s\n",
                  index + 1, contact.Name, contact.Email, contact.Phone)
            }
            fmt.Println(yellow("==============================================================="))

        case 3:
            // Eliminar un contacto
            fmt.Print(magenta("Ingrese el número del contacto que desea eliminar: "))
            var index int
            _, err := fmt.Scanf("%d\n", &index)
            if err != nil {
                fmt.Println(red("Error al leer el número:", err))
                continue
            }

            index-- // Ajustar a índice 0 basado

            if err := deleteContact(&contacts, index); err != nil {
                fmt.Println("Error al eliminar el contacto:", err)
                continue
            }

            // Guardar la lista actualizada de contactos
            if err := saveContactsFile(contacts); err != nil {
                fmt.Println(white("Error al guardar los contactos:", err))
            } else {
                fmt.Println(white("Contacto eliminado exitosamente."))
            }


        case 4:
            // Salir del programa
            fmt.Println(" ")
            fmt.Println(cyan("==================================================================="))
            fmt.Println(white("              Programador Sergio Alejandro Sopelana"))
            fmt.Println(cyan("==================================================================="))
            
            return

        default:
            fmt.Println(green("Opción no válida"))

        }
    }
 
 }

