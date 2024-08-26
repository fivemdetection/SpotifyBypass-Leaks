namespace Reverse
{
    partial class Login
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.components = new System.ComponentModel.Container();
            System.ComponentModel.ComponentResourceManager resources = new System.ComponentModel.ComponentResourceManager(typeof(Login));
            this.Exit = new Guna.UI2.WinForms.Guna2CircleButton();
            this.MinimizeButton = new Guna.UI2.WinForms.Guna2CircleButton();
            this.LoginButton = new Guna.UI2.WinForms.Guna2Button();
            this.guna2Elipse1 = new Guna.UI2.WinForms.Guna2Elipse(this.components);
            this.guna2DragControl1 = new Guna.UI2.WinForms.Guna2DragControl(this.components);
            this.EnterPass = new Guna.UI2.WinForms.Guna2TextBox();
            this.EnterUser = new Guna.UI2.WinForms.Guna2TextBox();
            this.guna2Button1 = new Guna.UI2.WinForms.Guna2Button();
            this.guna2Button2 = new Guna.UI2.WinForms.Guna2Button();
            this.SuspendLayout();
            // 
            // Exit
            // 
            this.Exit.BackColor = System.Drawing.Color.Transparent;
            this.Exit.CheckedState.Parent = this.Exit;
            this.Exit.CustomImages.Parent = this.Exit;
            this.Exit.FillColor = System.Drawing.Color.Transparent;
            this.Exit.Font = new System.Drawing.Font("Segoe UI", 9F);
            this.Exit.ForeColor = System.Drawing.Color.White;
            this.Exit.HoverState.Parent = this.Exit;
            this.Exit.Location = new System.Drawing.Point(769, 25);
            this.Exit.Name = "Exit";
            this.Exit.ShadowDecoration.Mode = Guna.UI2.WinForms.Enums.ShadowMode.Circle;
            this.Exit.ShadowDecoration.Parent = this.Exit;
            this.Exit.Size = new System.Drawing.Size(25, 25);
            this.Exit.TabIndex = 0;
            this.Exit.Click += new System.EventHandler(this.Exit_Click);
            // 
            // MinimizeButton
            // 
            this.MinimizeButton.BackColor = System.Drawing.Color.Transparent;
            this.MinimizeButton.CheckedState.Parent = this.MinimizeButton;
            this.MinimizeButton.CustomImages.Parent = this.MinimizeButton;
            this.MinimizeButton.FillColor = System.Drawing.Color.Transparent;
            this.MinimizeButton.Font = new System.Drawing.Font("Segoe UI", 9F);
            this.MinimizeButton.ForeColor = System.Drawing.Color.White;
            this.MinimizeButton.HoverState.Parent = this.MinimizeButton;
            this.MinimizeButton.Location = new System.Drawing.Point(738, 25);
            this.MinimizeButton.Name = "MinimizeButton";
            this.MinimizeButton.ShadowDecoration.Mode = Guna.UI2.WinForms.Enums.ShadowMode.Circle;
            this.MinimizeButton.ShadowDecoration.Parent = this.MinimizeButton;
            this.MinimizeButton.Size = new System.Drawing.Size(25, 25);
            this.MinimizeButton.TabIndex = 1;
            this.MinimizeButton.Click += new System.EventHandler(this.guna2CircleButton1_Click);
            // 
            // LoginButton
            // 
            this.LoginButton.Animated = true;
            this.LoginButton.BackColor = System.Drawing.Color.Transparent;
            this.LoginButton.BorderRadius = 15;
            this.LoginButton.CheckedState.Parent = this.LoginButton;
            this.LoginButton.CustomImages.Parent = this.LoginButton;
            this.LoginButton.FillColor = System.Drawing.Color.Transparent;
            this.LoginButton.Font = new System.Drawing.Font("Segoe UI", 9F);
            this.LoginButton.ForeColor = System.Drawing.Color.White;
            this.LoginButton.HoverState.Parent = this.LoginButton;
            this.LoginButton.Location = new System.Drawing.Point(214, 657);
            this.LoginButton.Name = "LoginButton";
            this.LoginButton.ShadowDecoration.Parent = this.LoginButton;
            this.LoginButton.Size = new System.Drawing.Size(314, 47);
            this.LoginButton.TabIndex = 2;
            this.LoginButton.Click += new System.EventHandler(this.LoginButton_Click);
            // 
            // guna2Elipse1
            // 
            this.guna2Elipse1.BorderRadius = 25;
            this.guna2Elipse1.TargetControl = this;
            // 
            // guna2DragControl1
            // 
            this.guna2DragControl1.TargetControl = this;
            // 
            // EnterPass
            // 
            this.EnterPass.BackColor = System.Drawing.Color.Transparent;
            this.EnterPass.BorderColor = System.Drawing.Color.Transparent;
            this.EnterPass.BorderRadius = 10;
            this.EnterPass.BorderThickness = 0;
            this.EnterPass.Cursor = System.Windows.Forms.Cursors.IBeam;
            this.EnterPass.DefaultText = "";
            this.EnterPass.DisabledState.BorderColor = System.Drawing.Color.FromArgb(((int)(((byte)(208)))), ((int)(((byte)(208)))), ((int)(((byte)(208)))));
            this.EnterPass.DisabledState.FillColor = System.Drawing.Color.FromArgb(((int)(((byte)(226)))), ((int)(((byte)(226)))), ((int)(((byte)(226)))));
            this.EnterPass.DisabledState.ForeColor = System.Drawing.Color.FromArgb(((int)(((byte)(138)))), ((int)(((byte)(138)))), ((int)(((byte)(138)))));
            this.EnterPass.DisabledState.Parent = this.EnterPass;
            this.EnterPass.DisabledState.PlaceholderForeColor = System.Drawing.Color.FromArgb(((int)(((byte)(138)))), ((int)(((byte)(138)))), ((int)(((byte)(138)))));
            this.EnterPass.FillColor = System.Drawing.Color.Transparent;
            this.EnterPass.FocusedState.BorderColor = System.Drawing.Color.FromArgb(((int)(((byte)(94)))), ((int)(((byte)(148)))), ((int)(((byte)(255)))));
            this.EnterPass.FocusedState.Parent = this.EnterPass;
            this.EnterPass.HoverState.BorderColor = System.Drawing.Color.FromArgb(((int)(((byte)(94)))), ((int)(((byte)(148)))), ((int)(((byte)(255)))));
            this.EnterPass.HoverState.Parent = this.EnterPass;
            this.EnterPass.IconLeft = ((System.Drawing.Image)(resources.GetObject("EnterPass.IconLeft")));
            this.EnterPass.Location = new System.Drawing.Point(207, 529);
            this.EnterPass.Name = "EnterPass";
            this.EnterPass.PasswordChar = '⦿';
            this.EnterPass.PlaceholderText = "Password";
            this.EnterPass.SelectedText = "";
            this.EnterPass.ShadowDecoration.Parent = this.EnterPass;
            this.EnterPass.Size = new System.Drawing.Size(321, 44);
            this.EnterPass.TabIndex = 3;
            // 
            // EnterUser
            // 
            this.EnterUser.BackColor = System.Drawing.Color.Transparent;
            this.EnterUser.BorderColor = System.Drawing.Color.Transparent;
            this.EnterUser.BorderRadius = 10;
            this.EnterUser.BorderThickness = 0;
            this.EnterUser.Cursor = System.Windows.Forms.Cursors.IBeam;
            this.EnterUser.DefaultText = "";
            this.EnterUser.DisabledState.BorderColor = System.Drawing.Color.FromArgb(((int)(((byte)(208)))), ((int)(((byte)(208)))), ((int)(((byte)(208)))));
            this.EnterUser.DisabledState.FillColor = System.Drawing.Color.FromArgb(((int)(((byte)(226)))), ((int)(((byte)(226)))), ((int)(((byte)(226)))));
            this.EnterUser.DisabledState.ForeColor = System.Drawing.Color.FromArgb(((int)(((byte)(138)))), ((int)(((byte)(138)))), ((int)(((byte)(138)))));
            this.EnterUser.DisabledState.Parent = this.EnterUser;
            this.EnterUser.DisabledState.PlaceholderForeColor = System.Drawing.Color.FromArgb(((int)(((byte)(138)))), ((int)(((byte)(138)))), ((int)(((byte)(138)))));
            this.EnterUser.FillColor = System.Drawing.Color.Transparent;
            this.EnterUser.FocusedState.BorderColor = System.Drawing.Color.FromArgb(((int)(((byte)(94)))), ((int)(((byte)(148)))), ((int)(((byte)(255)))));
            this.EnterUser.FocusedState.Parent = this.EnterUser;
            this.EnterUser.ForeColor = System.Drawing.Color.White;
            this.EnterUser.HoverState.BorderColor = System.Drawing.Color.FromArgb(((int)(((byte)(94)))), ((int)(((byte)(148)))), ((int)(((byte)(255)))));
            this.EnterUser.HoverState.Parent = this.EnterUser;
            this.EnterUser.IconLeft = ((System.Drawing.Image)(resources.GetObject("EnterUser.IconLeft")));
            this.EnterUser.Location = new System.Drawing.Point(207, 441);
            this.EnterUser.Name = "EnterUser";
            this.EnterUser.PasswordChar = '\0';
            this.EnterUser.PlaceholderText = "Username";
            this.EnterUser.SelectedText = "";
            this.EnterUser.ShadowDecoration.Parent = this.EnterUser;
            this.EnterUser.Size = new System.Drawing.Size(321, 39);
            this.EnterUser.TabIndex = 4;
            this.EnterUser.TextChanged += new System.EventHandler(this.EnterUser_TextChanged);
            // 
            // guna2Button1
            // 
            this.guna2Button1.Animated = true;
            this.guna2Button1.BackColor = System.Drawing.Color.Transparent;
            this.guna2Button1.BorderRadius = 15;
            this.guna2Button1.CheckedState.Parent = this.guna2Button1;
            this.guna2Button1.CustomImages.Parent = this.guna2Button1;
            this.guna2Button1.FillColor = System.Drawing.Color.Transparent;
            this.guna2Button1.Font = new System.Drawing.Font("Segoe UI", 9F);
            this.guna2Button1.ForeColor = System.Drawing.Color.White;
            this.guna2Button1.HoverState.Parent = this.guna2Button1;
            this.guna2Button1.Location = new System.Drawing.Point(692, 3);
            this.guna2Button1.Name = "guna2Button1";
            this.guna2Button1.ShadowDecoration.Parent = this.guna2Button1;
            this.guna2Button1.Size = new System.Drawing.Size(38, 30);
            this.guna2Button1.TabIndex = 5;
            this.guna2Button1.Text = "+";
            this.guna2Button1.Click += new System.EventHandler(this.guna2Button1_Click);
            // 
            // guna2Button2
            // 
            this.guna2Button2.Animated = true;
            this.guna2Button2.BackColor = System.Drawing.Color.Transparent;
            this.guna2Button2.BorderRadius = 15;
            this.guna2Button2.CheckedState.Parent = this.guna2Button2;
            this.guna2Button2.CustomImages.Parent = this.guna2Button2;
            this.guna2Button2.FillColor = System.Drawing.Color.Transparent;
            this.guna2Button2.Font = new System.Drawing.Font("Segoe UI", 9F);
            this.guna2Button2.ForeColor = System.Drawing.Color.White;
            this.guna2Button2.HoverState.Parent = this.guna2Button2;
            this.guna2Button2.Location = new System.Drawing.Point(648, 3);
            this.guna2Button2.Name = "guna2Button2";
            this.guna2Button2.ShadowDecoration.Parent = this.guna2Button2;
            this.guna2Button2.Size = new System.Drawing.Size(38, 30);
            this.guna2Button2.TabIndex = 6;
            this.guna2Button2.Text = "-";
            this.guna2Button2.Click += new System.EventHandler(this.guna2Button2_Click);
            // 
            // Login
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.BackColor = System.Drawing.Color.MidnightBlue;
            this.BackgroundImage = ((System.Drawing.Image)(resources.GetObject("$this.BackgroundImage")));
            this.ClientSize = new System.Drawing.Size(742, 907);
            this.Controls.Add(this.guna2Button2);
            this.Controls.Add(this.guna2Button1);
            this.Controls.Add(this.EnterUser);
            this.Controls.Add(this.EnterPass);
            this.Controls.Add(this.LoginButton);
            this.Controls.Add(this.MinimizeButton);
            this.Controls.Add(this.Exit);
            this.FormBorderStyle = System.Windows.Forms.FormBorderStyle.None;
            this.Icon = ((System.Drawing.Icon)(resources.GetObject("$this.Icon")));
            this.Name = "Login";
            this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen;
            this.Text = "Spotify";
            this.Load += new System.EventHandler(this.Form1_Load);
            this.ResumeLayout(false);

        }

        #endregion

        private Guna.UI2.WinForms.Guna2CircleButton Exit;
        private Guna.UI2.WinForms.Guna2CircleButton MinimizeButton;
        private Guna.UI2.WinForms.Guna2Button LoginButton;
        private Guna.UI2.WinForms.Guna2Elipse guna2Elipse1;
        private Guna.UI2.WinForms.Guna2DragControl guna2DragControl1;
        private Guna.UI2.WinForms.Guna2TextBox EnterPass;
        private Guna.UI2.WinForms.Guna2TextBox EnterUser;
        private Guna.UI2.WinForms.Guna2Button guna2Button1;
        private Guna.UI2.WinForms.Guna2Button guna2Button2;
    }
}

