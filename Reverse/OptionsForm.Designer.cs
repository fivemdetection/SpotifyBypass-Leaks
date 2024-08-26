namespace Reverse
{
    partial class OptionsForm
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
            System.ComponentModel.ComponentResourceManager resources = new System.ComponentModel.ComponentResourceManager(typeof(OptionsForm));
            this.guna2Button1 = new Guna.UI2.WinForms.Guna2Button();
            this.MinimizeButton = new Guna.UI2.WinForms.Guna2CircleButton();
            this.Exit = new Guna.UI2.WinForms.Guna2CircleButton();
            this.guna2DragControl1 = new Guna.UI2.WinForms.Guna2DragControl(this.components);
            this.SuspendLayout();
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
            this.guna2Button1.Location = new System.Drawing.Point(531, 549);
            this.guna2Button1.Name = "guna2Button1";
            this.guna2Button1.ShadowDecoration.Parent = this.guna2Button1;
            this.guna2Button1.Size = new System.Drawing.Size(185, 240);
            this.guna2Button1.TabIndex = 2;
            this.guna2Button1.Click += new System.EventHandler(this.guna2Button1_Click);
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
            this.MinimizeButton.Location = new System.Drawing.Point(1474, 1);
            this.MinimizeButton.Name = "MinimizeButton";
            this.MinimizeButton.ShadowDecoration.Mode = Guna.UI2.WinForms.Enums.ShadowMode.Circle;
            this.MinimizeButton.ShadowDecoration.Parent = this.MinimizeButton;
            this.MinimizeButton.Size = new System.Drawing.Size(25, 25);
            this.MinimizeButton.TabIndex = 5;
            this.MinimizeButton.Text = "Spotify";
            this.MinimizeButton.Click += new System.EventHandler(this.MinimizeButton_Click);
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
            this.Exit.Location = new System.Drawing.Point(1565, 1);
            this.Exit.Name = "Exit";
            this.Exit.ShadowDecoration.Mode = Guna.UI2.WinForms.Enums.ShadowMode.Circle;
            this.Exit.ShadowDecoration.Parent = this.Exit;
            this.Exit.Size = new System.Drawing.Size(25, 25);
            this.Exit.TabIndex = 4;
            this.Exit.Click += new System.EventHandler(this.Exit_Click);
            // 
            // guna2DragControl1
            // 
            this.guna2DragControl1.TargetControl = this;
            // 
            // OptionsForm
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.BackColor = System.Drawing.Color.MidnightBlue;
            this.BackgroundImage = ((System.Drawing.Image)(resources.GetObject("$this.BackgroundImage")));
            this.ClientSize = new System.Drawing.Size(1597, 895);
            this.Controls.Add(this.MinimizeButton);
            this.Controls.Add(this.Exit);
            this.Controls.Add(this.guna2Button1);
            this.FormBorderStyle = System.Windows.Forms.FormBorderStyle.None;
            this.Icon = ((System.Drawing.Icon)(resources.GetObject("$this.Icon")));
            this.Name = "OptionsForm";
            this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen;
            this.Text = "Spotify";
            this.Load += new System.EventHandler(this.OptionsForm_Load);
            this.ResumeLayout(false);

        }

        #endregion
        private Guna.UI2.WinForms.Guna2Button guna2Button1;
        private Guna.UI2.WinForms.Guna2CircleButton MinimizeButton;
        private Guna.UI2.WinForms.Guna2CircleButton Exit;
        private Guna.UI2.WinForms.Guna2DragControl guna2DragControl1;
    }
}