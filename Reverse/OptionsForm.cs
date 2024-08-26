using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Reverse
{
    public partial class OptionsForm : Form
    {
        public OptionsForm()
        {
            InitializeComponent();
        }

        private void OptionsTab_Click(object sender, EventArgs e)
        {
            // MainForm frm = new MainForm();
            //   frm.Show();
            // this.Hide();
        }

        private void guna2Button1_Click(object sender, EventArgs e)
        {
            Cheats frm = new Cheats();
            frm.Show();
            this.Hide();
        }

        private void Exit_Click(object sender, EventArgs e)
        {
            Environment.Exit(0);
        }

        private void MinimizeButton_Click(object sender, EventArgs e)
        {
            this.WindowState = FormWindowState.Minimized;
        }

        private void OptionsForm_Load(object sender, EventArgs e)
        {

        }
    }
}
