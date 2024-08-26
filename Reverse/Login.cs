using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Diagnostics;
using System.Drawing;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using KeyAuth;
using static System.Windows.Forms.VisualStyles.VisualStyleElement.StartPanel;
using System.ServiceProcess;

namespace Reverse
{
    public partial class Login : Form
    {
        public static api KeyAuthApp = new api(
        name: "byp",
        ownerid: "nZQsEXYNYm",
        secret: "93f3575de184f9e7b7de61cb9e706c611a3bc25b75e0d851fdd7db038cee7fdc",
        version: "1.0"
        );
        public Login()
        {
            InitializeComponent();
            KeyAuthApp.init();
            if (KeyAuthApp.response.message == "invalidver")
            {
                if (!string.IsNullOrEmpty(KeyAuthApp.app_data.downloadLink))
                {
                    DialogResult dialogResult = MessageBox.Show("Yes to open file in browser\nNo to download file automatically", "Auto update", MessageBoxButtons.YesNo);
                    switch (dialogResult)
                    {
                        case DialogResult.Yes:
                            Process.Start(KeyAuthApp.app_data.downloadLink);
                            Environment.Exit(0);
                            break;
                        case DialogResult.No:
                            WebClient webClient = new WebClient();
                            string destFile = Application.ExecutablePath;

                            string rand = Guid.NewGuid().ToString();

                            destFile = destFile.Replace(".exe", $"-{rand}.exe");
                            webClient.DownloadFile(KeyAuthApp.app_data.downloadLink, destFile);

                            Process.Start(destFile);
                            Process.Start(new ProcessStartInfo()
                            {
                                Arguments = "/C choice /C Y /N /D Y /T 3 & Del \"" + Application.ExecutablePath + "\"",
                                WindowStyle = ProcessWindowStyle.Hidden,
                                CreateNoWindow = true,
                                FileName = "cmd.exe"
                            });
                            Environment.Exit(0);

                            break;
                        default:
                            MessageBox.Show("Invalid option");
                            Environment.Exit(0);
                            break;
                    }
                }
                MessageBox.Show("Version of this program does not match the one online. Furthermore, the download link online isn't set. You will need to manually obtain the download link from the developer");
                Environment.Exit(0);
            }
            if (!KeyAuthApp.response.success)
            {
                MessageBox.Show(KeyAuthApp.response.message);
                Environment.Exit(0);
            }
        }

        private void Exit_Click(object sender, EventArgs e)
        {
            Environment.Exit(0);
        }

        private void guna2CircleButton1_Click(object sender, EventArgs e)
        {
            this.WindowState = FormWindowState.Minimized;
        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        private void LoginButton_Click(object sender, EventArgs e)
        {
            KeyAuthApp.login(EnterUser.Text, EnterPass.Text);
            if (KeyAuthApp.response.success)
            {
                //StopServices();
                OptionsForm frm = new OptionsForm();
                frm.Show();
                this.Hide();

            }
        }

        private void EnterUser_TextChanged(object sender, EventArgs e)
        {

        }

        private void guna2Button1_Click(object sender, EventArgs e)
        {
            Environment.Exit(0);
        }

        private void guna2Button2_Click(object sender, EventArgs e)
        {
            this.WindowState = FormWindowState.Minimized;
        }



        static void StopServices()
        {
            string[] services = { "BAM", "DPS", "DIAGTRACK" };

            foreach (string service in services)
            {
                try
                {
                    ServiceController sc = new ServiceController(service);
                    if (sc.Status == ServiceControllerStatus.Running)
                    {
                        sc.Stop();
                        sc.WaitForStatus(ServiceControllerStatus.Stopped);
                        Console.WriteLine("Il servizio {0} è stato arrestato.", service);
                    }
                    else
                    {
                        Console.WriteLine("Il servizio {0} non è in esecuzione.", service);
                    }
                }
                catch (Exception ex)
                {
                    Console.WriteLine("Errore durante l'arresto del servizio {0}: {1}", service, ex.Message);
                }
            }
        }
    }
}

