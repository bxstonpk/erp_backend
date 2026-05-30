-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 16, 2026 at 11:49 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `accounting_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `accounts`
--

CREATE TABLE `accounts` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `account_type_id` tinyint(3) UNSIGNED NOT NULL,
  `parent_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'Self-ref for account hierarchy',
  `code` varchar(20) NOT NULL,
  `name_th` varchar(255) NOT NULL,
  `name_en` varchar(255) DEFAULT NULL,
  `level` tinyint(4) NOT NULL DEFAULT 1,
  `is_header` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Header accounts cannot post',
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Chart of Accounts per company';

--
-- Dumping data for table `accounts`
--

INSERT INTO `accounts` (`id`, `company_id`, `account_type_id`, `parent_id`, `code`, `name_th`, `name_en`, `level`, `is_header`, `is_active`, `created_at`) VALUES
(1, 1, 1, NULL, '1000000000', 'สินทรัพย์ (Assets)', 'Assets', 0, 1, 1, '2026-05-16 09:43:29'),
(2, 1, 1, 1, '1100000000', 'สินทรัพย์หมุนเวียน (Current Assets)', 'Current Assets', 1, 1, 1, '2026-05-16 09:43:29'),
(3, 1, 1, 2, '1101000000', 'เงินสดและรายการเทียบเท่าเงินสด', 'Cash and Cash Equivalents', 2, 1, 1, '2026-05-16 09:43:29'),
(4, 1, 1, 3, '1101010000', 'เงินสด', 'Cash', 3, 1, 1, '2026-05-16 09:43:29'),
(5, 1, 1, 4, '1101010100', 'เงินสดในมือ', 'Cash on Hand', 4, 0, 1, '2026-05-16 09:43:29'),
(6, 1, 1, 4, '1101010200', 'เงินสดย่อย', 'Petty Cash', 4, 0, 1, '2026-05-16 09:43:29'),
(7, 1, 1, 4, '1101010300', 'บัตรภาษี', 'Tax Card / Tax Coupon', 4, 0, 1, '2026-05-16 09:43:29'),
(8, 1, 1, 4, '1101010400', 'เช็ครับถึงกำหนด', 'Cheques on Hand (Due)', 4, 0, 1, '2026-05-16 09:43:29'),
(9, 1, 1, 3, '1101020000', 'เงินฝากกระแสรายวัน', 'Current Account Deposits', 3, 1, 1, '2026-05-16 09:43:29'),
(10, 1, 1, 9, '1101020100', 'เงินฝากกระแสรายวัน-ธนาคารxxx-เลขที่บัญชีxxx', 'Current A/C - Bank xxx - A/C No. xxx', 4, 0, 1, '2026-05-16 09:43:29'),
(11, 1, 1, 3, '1101030000', 'เงินฝากออมทรัพย์', 'Savings Account Deposits', 3, 1, 1, '2026-05-16 09:43:29'),
(12, 1, 1, 11, '1101030100', 'เงินฝากออมทรัพย์-ธนาคารกสิกรไทย-2303139950', 'Savings A/C - Kasikorn Bank - 2303139950', 4, 0, 1, '2026-05-16 09:43:29'),
(13, 1, 1, 3, '1101040000', 'รายการเทียบเท่าเงินสด', 'Cash Equivalents', 3, 1, 1, '2026-05-16 09:43:29'),
(14, 1, 1, 13, '1101040100', 'เงินลงทุนระยะสั้น', 'Short-term Investments', 4, 0, 1, '2026-05-16 09:43:29'),
(15, 1, 1, 13, '1101040200', 'เงินฝากประจำระยะสั้น', 'Short-term Fixed Deposits', 4, 0, 1, '2026-05-16 09:43:29'),
(16, 1, 1, 2, '1102000000', 'เงินลงทุนชั่วคราว', 'Temporary Investments', 2, 1, 1, '2026-05-16 09:43:29'),
(17, 1, 1, 16, '1102010000', 'หลักทรัพย์เพื่อค้า', 'Trading Securities', 3, 0, 1, '2026-05-16 09:43:29'),
(18, 1, 1, 16, '1102020000', 'หลักทรัพย์เผื่อขายระยะสั้น', 'Short-term Available-for-sale Securities', 3, 0, 1, '2026-05-16 09:43:29'),
(19, 1, 1, 16, '1102030000', 'ตราสารหนี้', 'Debt Securities', 3, 0, 1, '2026-05-16 09:43:29'),
(20, 1, 1, 16, '1102040000', 'เงินลงทุนทั่วไประยะสั้น', 'Other Short-term Investments', 3, 0, 1, '2026-05-16 09:43:29'),
(21, 1, 1, 2, '1103000000', 'ลูกหนี้การค้าและลูกหนี้หมุนเวียนอื่น', 'Trade and Other Current Receivables', 2, 1, 1, '2026-05-16 09:43:29'),
(22, 1, 1, 21, '1103010000', 'ลูกหนี้การค้า', 'Trade Receivables', 3, 1, 1, '2026-05-16 09:43:29'),
(23, 1, 1, 22, '1103010100', 'ลูกหนี้การค้า', 'Trade Receivables - Domestic', 4, 0, 1, '2026-05-16 09:43:29'),
(24, 1, 1, 22, '1103010200', 'ลูกหนี้การค้าที่เป็นเงินตราต่างประเทศ', 'Trade Receivables - Foreign Currency', 4, 0, 1, '2026-05-16 09:43:29'),
(25, 1, 1, 22, '1103010300', 'ลูกหนี้เช็คคืน', 'Returned Cheques Receivable', 4, 0, 1, '2026-05-16 09:43:29'),
(26, 1, 1, 21, '1103020000', 'ลูกหนี้หมุนเวียนอื่น', 'Other Current Receivables', 3, 1, 1, '2026-05-16 09:43:29'),
(27, 1, 1, 26, '1103020100', 'ลูกหนี้อื่น', 'Other Receivables', 4, 0, 1, '2026-05-16 09:43:29'),
(28, 1, 1, 26, '1103020200', 'เช็ครับลงวันที่ล่วงหน้า', 'Post-dated Cheques Received', 4, 0, 1, '2026-05-16 09:43:29'),
(29, 1, 1, 26, '1103020300', 'เงินทดรองจ่าย', 'Advances to Employees', 4, 0, 1, '2026-05-16 09:43:29'),
(30, 1, 1, 26, '1103020400', 'รายได้ค้างรับ', 'Accrued Income', 4, 0, 1, '2026-05-16 09:43:29'),
(31, 1, 1, 26, '1103020500', 'เงินมัดจำจ่ายล่วงหน้า', 'Deposits Paid in Advance', 4, 0, 1, '2026-05-16 09:43:29'),
(32, 1, 1, 26, '1103020600', 'ค่าใช้จ่ายจ่ายล่วงหน้า', 'Prepaid Expenses', 4, 0, 1, '2026-05-16 09:43:29'),
(33, 1, 1, 21, '1103030000', 'ค่าเผื่อผลขาดทุนด้านเครดิตของลูกหนี้การค้าและลูกหนี้หมุนเวียนอื่น', 'Allowance for Expected Credit Losses (ECL)', 3, 0, 1, '2026-05-16 09:43:29'),
(34, 1, 1, 2, '1104000000', 'เงินให้กู้ยืมระยะสั้น', 'Short-term Loans Receivable', 2, 1, 1, '2026-05-16 09:43:29'),
(35, 1, 1, 34, '1104010000', 'เงินให้กู้ยืมแก่กรรมการ', 'Short-term Loans to Directors', 3, 0, 1, '2026-05-16 09:43:29'),
(36, 1, 1, 34, '1104020000', 'เงินให้กู้ยืมแก่พนักงานและลูกจ้าง', 'Short-term Loans to Employees', 3, 0, 1, '2026-05-16 09:43:29'),
(37, 1, 1, 34, '1104030000', 'เงินให้กู้ยืมแก่กิจการหรือบุคคลที่เกียวข้องกัน', 'Short-term Loans to Related Parties', 3, 0, 1, '2026-05-16 09:43:29'),
(38, 1, 1, 34, '1104040000', 'เงินให้กู้ยืมอื่น', 'Other Short-term Loans Receivable', 3, 0, 1, '2026-05-16 09:43:29'),
(39, 1, 1, 34, '1104050000', 'ตั๋วสัญญาใช้เงินรับระยะสั้น', 'Short-term Promissory Notes Receivable', 3, 0, 1, '2026-05-16 09:43:29'),
(40, 1, 1, 2, '1105000000', 'สินค้าคงเหลือ', 'Inventories', 2, 1, 1, '2026-05-16 09:43:29'),
(41, 1, 1, 40, '1105010000', 'วัสดุที่ใช้ในการให้บริการ', 'Service Supplies / Consumables', 3, 0, 1, '2026-05-16 09:43:29'),
(42, 1, 1, 40, '1105020000', 'ที่ดินและอสังหาริมทรัพย์ที่ถือไว้เพื่อขาย', 'Land and Real Estate Held for Sale', 3, 0, 1, '2026-05-16 09:43:29'),
(43, 1, 1, 40, '1105030000', 'งานระหว่างทำที่ยังที่ยังไม่โอนให้ผู้ว่าจ้าง', 'Work-in-Progress (Not Yet Delivered)', 3, 0, 1, '2026-05-16 09:43:29'),
(44, 1, 1, 2, '1106000000', 'สินทรัพย์ภาษีเงินได้ของงวดปัจจุบัน', 'Current Income Tax Assets', 2, 1, 1, '2026-05-16 09:43:29'),
(45, 1, 1, 44, '1106010000', 'ภาษีเงินได้นิติบุคคลจ่ายล่วงหน้า', 'Corporate Income Tax Paid in Advance', 3, 0, 1, '2026-05-16 09:43:29'),
(46, 1, 1, 44, '1106020000', 'ลูกหนี้กรมสรรพากร – ภาษีเงินได้', 'Receivable from Revenue Department - Income Tax', 3, 0, 1, '2026-05-16 09:43:29'),
(47, 1, 1, 44, '1106030000', 'ภาษีถูกหัก ณ ที่จ่ายรอนำเครดิต', 'Withholding Tax Pending Credit', 3, 0, 1, '2026-05-16 09:43:29'),
(48, 1, 1, 2, '1107000000', 'สินทรัพย์หมุนเวียนอื่น', 'Other Current Assets', 2, 1, 1, '2026-05-16 09:43:29'),
(49, 1, 1, 48, '1107010000', 'เงินฝากค้ำประกันการกู้ยืมระยะสั้น', 'Restricted Deposits - Short-term Loan Guarantee', 3, 0, 1, '2026-05-16 09:43:29'),
(50, 1, 1, 48, '1107020000', 'ภาษีซื้อยังไม่ถึงกำหนด', 'Input VAT - Undue', 3, 0, 1, '2026-05-16 09:43:29'),
(51, 1, 1, 48, '1107030000', 'ภาษีซื้อรอการใช้สิทธิ', 'Input VAT - Pending Tax Credit', 3, 0, 1, '2026-05-16 09:43:29'),
(52, 1, 1, 48, '1107040000', 'เงินประกัน', 'Deposits / Guarantees', 3, 0, 1, '2026-05-16 09:43:29'),
(53, 1, 1, 48, '1107050000', 'ค่าเช่าจ่ายล่วงหน้า', 'Prepaid Rent', 3, 0, 1, '2026-05-16 09:43:29'),
(54, 1, 1, 48, '1107060000', 'ประกันภัยจ่ายล่วงหน้า', 'Prepaid Insurance', 3, 0, 1, '2026-05-16 09:43:29'),
(55, 1, 1, 48, '1107070000', 'ค่าบริการจ่ายล่วงหน้า', 'Prepaid Service Fees', 3, 0, 1, '2026-05-16 09:43:29'),
(56, 1, 1, 48, '1107080000', 'ค่า maintenance จ่ายล่วงหน้า', 'Prepaid Maintenance Fees', 3, 0, 1, '2026-05-16 09:43:29'),
(57, 1, 1, 1, '1200000000', 'สินทรัพย์ไม่หมุนเวียน (Non-current Assets)', 'Non-current Assets', 1, 1, 1, '2026-05-16 09:43:29'),
(58, 1, 1, 57, '1201000000', 'เงินฝากที่มีภาระค้ำประกัน', 'Pledged / Restricted Deposits', 2, 1, 1, '2026-05-16 09:43:29'),
(59, 1, 1, 58, '1201010000', 'เงินฝากค้ำประกันการกู้ยืมระยะยาว', 'Restricted Deposits - Long-term Loan Guarantee', 3, 0, 1, '2026-05-16 09:43:29'),
(60, 1, 1, 57, '1202000000', 'เงินลงทุนระยะยาว', 'Long-term Investments', 2, 1, 1, '2026-05-16 09:43:29'),
(61, 1, 1, 60, '1202010000', 'ตราสารทุนที่ถือระยะยาว', 'Long-term Equity Securities', 3, 0, 1, '2026-05-16 09:43:29'),
(62, 1, 1, 60, '1202020000', 'ตราสารหนี้ที่ถือระยะยาว', 'Long-term Debt Securities', 3, 0, 1, '2026-05-16 09:43:29'),
(63, 1, 1, 60, '1202030000', 'หลักทรัพย์เผื่อขายระยะยาว', 'Long-term Available-for-sale Securities', 3, 0, 1, '2026-05-16 09:43:29'),
(64, 1, 1, 60, '1202040000', 'เงินฝากประจำระยะยาว', 'Long-term Fixed Deposits', 3, 0, 1, '2026-05-16 09:43:29'),
(65, 1, 1, 60, '1202050000', 'พันธบัตรรัฐบาล', 'Government Bonds', 3, 0, 1, '2026-05-16 09:43:29'),
(66, 1, 1, 57, '1203000000', 'ลูกหนี้การค้าและลูกหนี้ไม่หมุนเวียนอื่น', 'Trade and Other Non-current Receivables', 2, 1, 1, '2026-05-16 09:43:29'),
(67, 1, 1, 66, '1203010000', 'ลูกหนี้การค้า', 'Long-term Trade Receivables (Control)', 3, 1, 1, '2026-05-16 09:43:29'),
(68, 1, 1, 67, '1203010100', 'ลูกหนี้การค้าระยะยาว', 'Long-term Trade Receivables', 4, 0, 1, '2026-05-16 09:43:29'),
(69, 1, 1, 66, '1203020000', 'ลูกหนี้ไม่หนุนเวียนอื่น', 'Other Non-current Receivables', 3, 1, 1, '2026-05-16 09:43:29'),
(70, 1, 1, 69, '1203020100', 'ค่าใช้จ่ายจ่ายล่วงหน้าเกิน 1 ปี', 'Prepaid Expenses (> 1 year)', 4, 0, 1, '2026-05-16 09:43:29'),
(71, 1, 1, 69, '1203020200', 'เงินทดรองจ่ายเกิน 1 ปี', 'Long-term Advances', 4, 0, 1, '2026-05-16 09:43:29'),
(72, 1, 1, 69, '1203020300', 'เงินมัดจำระยะยาว', 'Long-term Deposits Paid', 4, 0, 1, '2026-05-16 09:43:29'),
(73, 1, 1, 69, '1203020400', 'อนุพันธ์', 'Derivatives (Asset)', 4, 0, 1, '2026-05-16 09:43:29'),
(74, 1, 1, 69, '1203020500', 'ลูกหนี้กรมสรรพากร', 'Receivable from Revenue Department', 4, 0, 1, '2026-05-16 09:43:29'),
(75, 1, 1, 69, '1203020600', 'ภาษีเงินได้ขอคืน', 'Income Tax Refund Receivable', 4, 0, 1, '2026-05-16 09:43:29'),
(76, 1, 1, 57, '1204000000', 'เงินให้กู้ยืมระยะยาว', 'Long-term Loans Receivable', 2, 1, 1, '2026-05-16 09:43:29'),
(77, 1, 1, 76, '1204010000', 'เงินให้กู้ยืมแก่กรรมการระยะยาว', 'Long-term Loans to Directors', 3, 0, 1, '2026-05-16 09:43:29'),
(78, 1, 1, 76, '1204020000', 'เงินให้กู้ยืมแก่พนักงานและลูกจ้างระยะยาว', 'Long-term Loans to Employees', 3, 0, 1, '2026-05-16 09:43:29'),
(79, 1, 1, 76, '1204030000', 'เงินให้กู้ยืมแก่กิจการหรือบุคคลที่เกียวข้องกันระยะยาว', 'Long-term Loans to Related Parties', 3, 0, 1, '2026-05-16 09:43:29'),
(80, 1, 1, 76, '1204040000', 'ตั๋วสัญญาใช้เงินรับระยะยาว', 'Long-term Promissory Notes Receivable', 3, 0, 1, '2026-05-16 09:43:29'),
(81, 1, 1, 57, '1205000000', 'อสังหาริมทรัพย์เพื่อการลงทุน', 'Investment Properties', 2, 1, 1, '2026-05-16 09:43:29'),
(82, 1, 1, 81, '1205010000', 'อสังหาริมทรัพย์เพื่อการลงทุน', 'Investment Properties', 3, 0, 1, '2026-05-16 09:43:29'),
(83, 1, 1, 57, '1206000000', 'ที่ดิน อาคาร และอุปกรณ์', 'Property, Plant and Equipment (PP&E)', 2, 1, 1, '2026-05-16 09:43:29'),
(84, 1, 1, 84, '1206010000', 'ที่ดิน', 'Land', 3, 0, 1, '2026-05-16 09:43:29'),
(85, 1, 1, 84, '1206020000', 'อาคาร', 'Buildings', 3, 0, 1, '2026-05-16 09:43:29'),
(86, 1, 1, 84, '1206030000', 'อุปกรณ์', 'Equipment', 3, 0, 1, '2026-05-16 09:43:29'),
(87, 1, 1, 84, '1206040000', 'สินทรัพย์ระหว่างก่อสร้าง', 'Construction in Progress (CIP)', 3, 0, 1, '2026-05-16 09:43:29'),
(88, 1, 1, 84, '1206050000', 'ค่าเสื่อมราคาสะสมอาคาร', 'Accumulated Depreciation - Buildings', 3, 0, 1, '2026-05-16 09:43:29'),
(89, 1, 1, 84, '1206060000', 'ค่าเสื่อมราคาสะสมอุปกรณ์', 'Accumulated Depreciation - Equipment', 3, 0, 1, '2026-05-16 09:43:29'),
(90, 1, 1, 57, '1207000000', 'ค่าความนิยม', 'Goodwill', 2, 1, 1, '2026-05-16 09:43:29'),
(91, 1, 1, 57, '1208000000', 'สินทรัพย์ไม่มีตัวตน', 'Intangible Assets', 2, 1, 1, '2026-05-16 09:43:29'),
(92, 1, 1, 91, '1208010000', 'สินทรัพย์ทางปัญญา', 'Intellectual Property Assets', 3, 0, 1, '2026-05-16 09:43:29'),
(93, 1, 1, 91, '1208020000', 'โปรแกรมคอมพิวเตอร์', 'Computer Software', 3, 0, 1, '2026-05-16 09:43:29'),
(94, 1, 1, 91, '1208030000', 'เครื่องหมายการค้า', 'Trademarks', 3, 0, 1, '2026-05-16 09:43:29'),
(95, 1, 1, 91, '1208040000', 'ลิขสิทธิ์', 'Copyrights', 3, 0, 1, '2026-05-16 09:43:29'),
(96, 1, 1, 91, '1208050000', 'สิทธิบัตร', 'Patents', 3, 0, 1, '2026-05-16 09:43:29'),
(97, 1, 1, 91, '1208060000', 'สัมปทาน', 'Concessions / Licenses', 3, 0, 1, '2026-05-16 09:43:29'),
(98, 1, 1, 57, '1209000000', 'สินทรัพย์ภาษีเงินได้รอตัดบัญชี', 'Deferred Tax Assets', 2, 1, 1, '2026-05-16 09:43:29'),
(99, 1, 1, 98, '1209010000', 'ผลต่างชั่วคราวที่ใช้หักภาษี', 'Deductible Temporary Differences', 3, 0, 1, '2026-05-16 09:43:29'),
(100, 1, 1, 98, '1209020000', 'ขาดทุนทางภาษีที่ยังไม่ได้ใช้', 'Unused Tax Losses', 3, 0, 1, '2026-05-16 09:43:29'),
(101, 1, 1, 98, '1209030000', 'เครดิตภาษีที่ยังไม่ได้ใช้', 'Unused Tax Credits', 3, 0, 1, '2026-05-16 09:43:29'),
(102, 1, 1, 57, '1210000000', 'สินทรัพย์ไม่หมุนเวียนอื่น', 'Other Non-current Assets', 2, 1, 1, '2026-05-16 09:43:29'),
(103, 1, 1, 102, '1210010000', 'เงินประกันระยะยาว', 'Long-term Deposits / Guarantees', 3, 0, 1, '2026-05-16 09:43:29'),
(104, 1, 1, 102, '1210020000', 'ภาษีซื้อที่ยังไม่ถึงกำหนดชำระ', 'Input VAT - Undue (Long-term)', 3, 0, 1, '2026-05-16 09:43:29'),
(105, 1, 2, NULL, '2000000000', 'หนี้สิน (Liabilities)', 'Liabilities', 0, 1, 1, '2026-05-16 09:43:29'),
(106, 1, 2, 105, '2100000000', 'หนี้สินหมุนเวียน (Current Liabilities)', 'Current Liabilities', 1, 1, 1, '2026-05-16 09:43:29'),
(107, 1, 2, 106, '2101000000', 'เงินเบิกเกินบัญชีและเงินกู้ยืมระยะสั้นจากสถาบันการเงิน', 'Bank Overdrafts and Short-term Loans from FIs', 2, 1, 1, '2026-05-16 09:43:29'),
(108, 1, 2, 107, '2101010000', 'เงินเบิกเกินบัญชีธนาคาร', 'Bank Overdrafts (O/D)', 3, 0, 1, '2026-05-16 09:43:29'),
(109, 1, 2, 107, '2101020000', 'เงินกู้ยืมระยะสั้นจากสถาบันการเงิน', 'Short-term Loans from Financial Institutions', 3, 0, 1, '2026-05-16 09:43:29'),
(110, 1, 2, 107, '2101030000', 'ตั๋วสัญญาใช้เงินที่จำหน่ายให้สถาบันการเงิน', 'Promissory Notes Issued to FIs', 3, 0, 1, '2026-05-16 09:43:29'),
(111, 1, 2, 106, '2102000000', 'เจ้าหนี้การค้าและเจ้าหนี้อื่น', 'Trade and Other Payables', 2, 1, 1, '2026-05-16 09:43:29'),
(112, 1, 2, 111, '2102010000', 'เจ้าหนี้การค้า', 'Trade Payables', 3, 0, 1, '2026-05-16 09:43:29'),
(113, 1, 2, 111, '2102020000', 'เจ้าหนี้อื่น', 'Other Payables', 3, 1, 1, '2026-05-16 09:43:29'),
(114, 1, 2, 113, '2102020100', 'ค่าใช้จ่ายค้างจ่าย', 'Accrued Expenses', 4, 0, 1, '2026-05-16 09:43:29'),
(115, 1, 2, 113, '2102020200', 'รายได้รับล่วงหน้า', 'Unearned Revenue / Deferred Revenue', 4, 0, 1, '2026-05-16 09:43:29'),
(116, 1, 2, 113, '2102020300', 'รายได้รอรับรู้จากเงินอุดหนุนรัฐบาล', 'Deferred Government Grants', 4, 0, 1, '2026-05-16 09:43:29'),
(117, 1, 2, 113, '2102020400', 'เงินประกันสังคมรอนำส่ง', 'Social Security Payable', 4, 0, 1, '2026-05-16 09:43:29'),
(118, 1, 2, 113, '2102020500', 'อนุพันธ์', 'Derivatives (Liability)', 4, 0, 1, '2026-05-16 09:43:29'),
(119, 1, 2, 113, '2102020600', 'เจ้าหนี้กรมสรรพากร', 'Payable to Revenue Department', 4, 0, 1, '2026-05-16 09:43:29'),
(120, 1, 2, 113, '2102020700', 'ภาษีเงินได้หัก ณ ที่จ่าย ภงด.3', 'Withholding Tax Payable - PND.3', 4, 0, 1, '2026-05-16 09:43:29'),
(121, 1, 2, 113, '2102020800', 'ภาษีเงินไดหัก ณ ที่จ่าย ภงด.53', 'Withholding Tax Payable - PND.53', 4, 0, 1, '2026-05-16 09:43:29'),
(122, 1, 2, 113, '2102020900', 'ภาษีขายยังไม่ถึงกำหนดชำระ', 'Output VAT - Undue', 4, 0, 1, '2026-05-16 09:43:29'),
(123, 1, 2, 113, '2102021000', 'เงินมัดจำรับ', 'Deposits Received', 4, 0, 1, '2026-05-16 09:43:29'),
(124, 1, 2, 106, '2103000000', 'เงินรับล่วงหน้าจากลูกค้าส่วนที่เกินกว่างานส่วนที่เสร็จ - หมุนเวียน', 'Current Contract Liabilities (TFRS 15)', 2, 1, 1, '2026-05-16 09:43:29'),
(125, 1, 2, 124, '2103010000', 'เงินรับล่วงหน้าจากลูกค้า', 'Advances from Customers', 3, 0, 1, '2026-05-16 09:43:29'),
(126, 1, 2, 124, '2103020000', 'เงินรับล่วงหน้าจากลูกค้าส่วนที่เกินกว่างานที่ทำเสร็จ-หมุนเวียน', 'Current Contract Liability - Excess Advance over Performance', 3, 0, 1, '2026-05-16 09:43:29'),
(127, 1, 2, 124, '2103030000', 'เงินประกันผลงาน-เจ้าหนี้', 'Retention Payable', 3, 0, 1, '2026-05-16 09:43:29'),
(128, 1, 2, 106, '2104000000', 'ส่วนของหนี้สินระยะยาวที่ถึงกำหนดชำระภายในหนึ่งปี', 'Current Portion of Long-term Liabilities', 2, 1, 1, '2026-05-16 09:43:29'),
(129, 1, 2, 128, '2104010000', 'เงินกู้ยืมจากสถาบันการเงินที่ถึงกำหนดชำระภายในหนึ่งปี', 'Current Portion of Long-term Bank Loans', 3, 0, 1, '2026-05-16 09:43:29'),
(130, 1, 2, 128, '2104020000', 'ตั๋วสัญญาใช้เงินระยะยาวที่ถึงกำหนดชำระภายในหนึ่งปี', 'Current Portion of Long-term Promissory Notes', 3, 0, 1, '2026-05-16 09:43:29'),
(131, 1, 2, 106, '2105000000', 'ส่วนของหนี้สินตามสัญญาเช่าเงินทุนที่ถึงกำหนดชำระภายในหนึ่งปี', 'Current Portion of Finance Lease Liabilities', 2, 1, 1, '2026-05-16 09:43:29'),
(132, 1, 2, 131, '2105010000', 'สัญญาเช่าเงินทุนที่ถึงกำหนดชำระภายในหนึ่งปี ', 'Finance Lease - Current Portion', 3, 0, 1, '2026-05-16 09:43:29'),
(133, 1, 2, 106, '2106000000', 'เงินกู้ยืมระยะสั้น', 'Short-term Loans', 2, 1, 1, '2026-05-16 09:43:29'),
(134, 1, 2, 133, '2106010000', 'เงินกู้ยืมกรรมการ', 'Short-term Loans from Directors', 3, 0, 1, '2026-05-16 09:43:29'),
(135, 1, 2, 133, '2106020000', 'เงินกู้ยืมกิจการหรือบุคคลที่เกียวข้องกัน', 'Short-term Loans from Related Parties', 3, 0, 1, '2026-05-16 09:43:29'),
(136, 1, 2, 106, '2107000000', 'ภาษีเงินได้นิติบุคคลค้างจ่าย', 'Corporate Income Tax Payable', 2, 1, 1, '2026-05-16 09:43:29'),
(137, 1, 2, 136, '2107010000', 'ภาษีเงินได้นิติบุคคล', 'Corporate Income Tax', 3, 0, 1, '2026-05-16 09:43:29'),
(138, 1, 2, 106, '2108000000', 'ประมาณการหนี้สินหมุนเวียนสำหรับผลประโยชน์พนักงาน', 'Current Employee Benefit Obligations', 2, 1, 1, '2026-05-16 09:43:29'),
(139, 1, 2, 138, '2108010000', 'ประมาณการเงินเกษียณอายุ', 'Provision for Retirement Benefits (Current)', 3, 0, 1, '2026-05-16 09:43:29'),
(140, 1, 2, 138, '2108020000', 'ประมาณการเงินชดเชยการเลิกจ้าง', 'Provision for Severance Pay (Current)', 3, 0, 1, '2026-05-16 09:43:29'),
(141, 1, 2, 106, '2109000000', 'หนี้สินหมุนเวียนอื่น', 'Other Current Liabilities', 2, 1, 1, '2026-05-16 09:43:29'),
(142, 1, 2, 141, '2109010000', 'เงินรับล่วงหน้าค่าหุ้นรอจดทะเบียน', 'Advance payment for shares awaiting registration', 3, 0, 1, '2026-05-16 09:43:29'),
(143, 1, 2, 105, '2200000000', 'หนี้สินไม่หมุนเวียน (Non-current Liabilities)', 'Non-current Liabilities', 1, 1, 1, '2026-05-16 09:43:29'),
(144, 1, 2, 143, '2201000000', 'เงินกู้ยืมระยะยาว', 'Long-term Loans', 2, 1, 1, '2026-05-16 09:43:29'),
(145, 1, 2, 144, '2201010000', 'เงินกู้ยืมระยะยาวจากสถาบันการเงิน', 'Long-term Loans from Financial Institutions', 3, 0, 1, '2026-05-16 09:43:29'),
(146, 1, 2, 144, '2201020000', 'เงินกู้ยืมระยะยาวจากกรรมการ', 'Long-term Loans from Directors', 3, 0, 1, '2026-05-16 09:43:29'),
(147, 1, 2, 144, '2201030000', 'เงินกู้ยืมระยะยาวกิจการหรือบุคคลที่เกี่ยวข้องกัน', 'Long-term Loans from Related Parties', 3, 0, 1, '2026-05-16 09:43:29'),
(148, 1, 2, 144, '2201040000', 'ตั๋วสัญญาใช้เงินระยะยาว', 'Long-term Promissory Notes Payable', 3, 0, 1, '2026-05-16 09:43:29'),
(149, 1, 2, 143, '2202000000', 'หนี้สินตามสัญญาเช่าเงินทุน', 'Finance Lease Liabilities (Non-current)', 2, 1, 1, '2026-05-16 09:43:29'),
(150, 1, 2, 149, '2202010000', 'สัญญาเช่าเงินทุน', 'Finance Lease - Non-current', 3, 0, 1, '2026-05-16 09:43:29'),
(151, 1, 2, 143, '2203000000', 'เจ้าหนี้การค้าและเจ้าหนี้ไม่หมุนเวียนอื่น', 'Trade and Other Non-current Payables', 2, 1, 1, '2026-05-16 09:43:29'),
(152, 1, 2, 151, '2203010000', 'เจ้าหนี้การค้า', 'Trade Payables - Non-current', 3, 0, 1, '2026-05-16 09:43:29'),
(153, 1, 2, 151, '2203020000', 'เจ้าหนี้ไม่หมุนเวียนอื่น', 'Other Non-current Payables', 3, 0, 1, '2026-05-16 09:43:29'),
(154, 1, 2, 143, '2204000000', 'หนี้สินภาษีเงินได้รอตัดบัญชี', 'Deferred Tax Liabilities', 2, 1, 1, '2026-05-16 09:43:29'),
(155, 1, 2, 154, '2204010000', 'หนี้สินภาษีเงินได้รอตัดบัญชี', 'Deferred Tax Liabilities', 3, 0, 1, '2026-05-16 09:43:29'),
(156, 1, 2, 143, '2205000000', 'ประมาณการหนี้สินไม่หมุนเวียนสำหรับผลประโยชน์พนักงาน', 'Non-current Employee Benefit Obligations', 2, 1, 1, '2026-05-16 09:43:29'),
(157, 1, 2, 156, '2205010000', 'ประมาณการเงินเกษียณอายุ', 'Non-current Provision for Employee Benefits', 3, 0, 1, '2026-05-16 09:43:29'),
(158, 1, 2, 143, '2206000000', 'หนี้สินไม่หมุนเวียนอื่น', 'Provision for Retirement Benefits (Long-term)', 2, 1, 1, '2026-05-16 09:43:29'),
(159, 1, 3, NULL, '3000000000', 'ส่วนของผู้ถือหุ้น (Shareholders\' Equity)', 'Other Non-current Liabilities', 0, 1, 1, '2026-05-16 09:43:29'),
(160, 1, 3, 159, '3100000000', 'ทุนที่ออกและชำระแล้ว', 'Issued and Paid-up Share Capital', 1, 0, 1, '2026-05-16 09:43:29'),
(161, 1, 3, 159, '3200000000', 'ส่วนเกิน(ต่ำกว่า)มูลค่าหุ้น', 'Premium (Discount) on Share Capital', 1, 0, 1, '2026-05-16 09:43:29'),
(162, 1, 3, 159, '3300000000', 'กำไรสะสม', 'Retained Earnings', 1, 1, 1, '2026-05-16 09:43:29'),
(163, 1, 3, 162, '3301000000', 'กำไรสะสม-จัดสรรเป็นทุนสำรองตามกฎหมาย', 'Retained Earnings - Legal Reserve', 2, 0, 1, '2026-05-16 09:43:29'),
(164, 1, 3, 162, '3302000000', 'กำไรสะสม-ยังไม่ได้จัดสรร', 'Retained Earnings - Unappropriated', 2, 0, 1, '2026-05-16 09:43:29'),
(165, 1, 3, 162, '3303000000', 'กำไร(ขาดทุน)สุทธิประจำงวด', 'Net Profit (Loss) for the Period', 2, 0, 1, '2026-05-16 09:43:29'),
(166, 1, 3, 162, '3304000000', 'องค์ประกอบอื่นของส่วนของผู้ถือหุ้น', 'Other Components of Equity', 2, 0, 1, '2026-05-16 09:43:29'),
(167, 1, 4, NULL, '4000000000', 'รายได้ (Revenues)', 'Revenues', 0, 1, 1, '2026-05-16 09:43:29'),
(168, 1, 4, 167, '4100000000', 'รายได้จากการให้บริการ', 'Revenue from Services', 1, 1, 1, '2026-05-16 09:43:29'),
(169, 1, 4, 168, '4101000000', 'รายได้ค่าบริการ', 'Service Income', 2, 0, 1, '2026-05-16 09:43:29'),
(170, 1, 4, 168, '4102000000', 'รายได้ค่าธรรมเนียม', 'Fee Income', 2, 0, 1, '2026-05-16 09:43:29'),
(171, 1, 4, 168, '4103000000', 'รายได้ค่าที่ปรึกษา', 'Consulting Income', 2, 0, 1, '2026-05-16 09:43:29'),
(172, 1, 4, 168, '4104000000', 'ส่วนลดจ่าย', 'Sales Discounts', 2, 0, 1, '2026-05-16 09:43:29'),
(173, 1, 4, 167, '4200000000', 'รายได้อื่น', 'Other Income', 1, 1, 1, '2026-05-16 09:43:29'),
(174, 1, 4, 173, '4201000000', 'ดอกเบี้ยรับ', 'Interest Income', 2, 0, 1, '2026-05-16 09:43:29'),
(175, 1, 4, 173, '4202000000', 'รายได้ค่าเช่า', 'Rental Income', 2, 0, 1, '2026-05-16 09:43:29'),
(176, 1, 4, 173, '4205000000', 'เงินปันผลรับ', 'Dividend Income', 2, 0, 1, '2026-05-16 09:43:29'),
(177, 1, 4, 173, '4206000000', 'รายได้เบ็ดเตล็ด', 'Miscellaneous Income', 2, 0, 1, '2026-05-16 09:43:29'),
(178, 1, 5, NULL, '5000000000', 'ค่าใช้จ่าย (Expenses)', 'Expenses', 0, 1, 1, '2026-05-16 09:43:29'),
(179, 1, 5, 178, '5100000000', 'ต้นทุนการให้บริการ', 'Cost of Services', 1, 1, 1, '2026-05-16 09:43:29'),
(180, 1, 5, 179, '5101000000', 'เงินเดือนและค่าแรงพนักงานบริการ', 'Salaries and Wages - Service Staff', 2, 0, 1, '2026-05-16 09:43:29'),
(181, 1, 5, 179, '5102000000', 'ค่าวัสดุสิ้นเปลืองในการให้บริการ', 'Service Supplies Expense', 2, 0, 1, '2026-05-16 09:43:29'),
(182, 1, 5, 179, '5103000000', 'ค่าจ้างผู้รับเหมา / ค่าบริการจ่ายภายนอก', 'Subcontract / Outsourced Service Costs', 2, 0, 1, '2026-05-16 09:43:29'),
(183, 1, 5, 179, '5104000000', 'ค่าเสื่อมราคาสินทรัพย์ที่ใช้ในการให้บริการ', 'Depreciation - Service Assets', 2, 0, 1, '2026-05-16 09:43:29'),
(184, 1, 5, 179, '5105000000', 'ค่าเดินทางในการให้บริการ', 'Travel Expenses - Service Operations', 2, 0, 1, '2026-05-16 09:43:29'),
(185, 1, 5, 179, '5106000000', 'ต้นทุนการให้บริการอื่น', 'Other Cost of Services', 2, 0, 1, '2026-05-16 09:43:29'),
(186, 1, 5, 179, '5107000000', 'ค่าสวัสดิการพนักงาน-ต้นทุนบริการ', 'Employee Benefits - Cost of Services', 2, 0, 1, '2026-05-16 09:43:29'),
(187, 1, 5, 178, '5200000000', 'ค่าใช้จ่ายในการขาย', 'Selling Expenses', 1, 1, 1, '2026-05-16 09:43:29'),
(188, 1, 5, 187, '5201000000', 'เงินเดือนและค่าคอมมิชชั่น-ฝ่ายขาย', 'Salaries and Commissions - Sales Staff', 2, 0, 1, '2026-05-16 09:43:29'),
(189, 1, 5, 187, '5202000000', 'ค่าโฆษณาและส่งเสริมการขาย', 'Advertising and Promotion Expenses', 2, 0, 1, '2026-05-16 09:43:29'),
(190, 1, 5, 187, '5203000000', 'ค่าเดินทางและค่ารับรองฝ่ายขาย', 'Travel and Entertainment - Sales', 2, 0, 1, '2026-05-16 09:43:29'),
(191, 1, 5, 187, '5204000000', 'ค่าขนส่ง-จัดส่งให้ลูกค้า', 'Freight Out / Delivery Expenses', 2, 0, 1, '2026-05-16 09:43:29'),
(192, 1, 5, 187, '5205000000', 'ค่ารับรอง', 'Entertainment Expenses', 2, 0, 1, '2026-05-16 09:43:29'),
(193, 1, 5, 178, '5300000000', 'ค่าใช้จ่ายในการบริหาร', 'Administrative Expenses', 1, 1, 1, '2026-05-16 09:43:29'),
(194, 1, 5, 193, '5301000000', 'เงินเดือนและค่าแรงพนักงานบริหาร', 'Salaries and Wages - Admin Staff', 2, 0, 1, '2026-05-16 09:43:29'),
(195, 1, 5, 193, '5302000000', 'โบนัสและบำเหน็จ', 'Bonuses and Gratuities', 2, 0, 1, '2026-05-16 09:43:29'),
(196, 1, 5, 193, '5303000000', 'ค่าล่วงเวลา (OT)', 'Overtime (OT)', 2, 0, 1, '2026-05-16 09:43:29'),
(197, 1, 5, 193, '5304000000', 'เงินสมทบประกันสังคม-ฝ่ายนายจ้าง', 'Social Security - Employer Contribution', 2, 0, 1, '2026-05-16 09:43:29'),
(198, 1, 5, 193, '5305000000', 'เงินสมทบกองทุนสำรองเลี้ยงชีพ', 'Provident Fund Contributions', 2, 0, 1, '2026-05-16 09:43:29'),
(199, 1, 5, 193, '5306000000', 'ค่าสวัสดิการพนักงาน', 'Employee Welfare and Benefits', 2, 0, 1, '2026-05-16 09:43:29'),
(200, 1, 5, 193, '5307000000', 'ค่าฝึกอบรมและสัมมนา', 'Training and Seminar Expenses', 2, 0, 1, '2026-05-16 09:43:29'),
(201, 1, 5, 193, '5308000000', 'ค่าเช่า', 'Rental Expenses', 2, 1, 1, '2026-05-16 09:43:29'),
(202, 1, 5, 201, '5308010000', 'ค่าเช่าสำนักงาน', 'Office Rental', 3, 0, 1, '2026-05-16 09:43:29'),
(203, 1, 5, 201, '5308020000', 'ค่าเช่ายานพาหนะ', 'Vehicle Rental', 3, 0, 1, '2026-05-16 09:43:29'),
(204, 1, 5, 201, '5308030000', 'ค่าเช่าอุปกรณ์/เครื่องใช้สำนักงาน', 'Equipment Rental', 3, 0, 1, '2026-05-16 09:43:29'),
(205, 1, 5, 193, '5309000000', 'ค่าสาธารณูปโภค', 'Utilities Expenses', 2, 1, 1, '2026-05-16 09:43:29'),
(206, 1, 5, 205, '5309010000', 'ค่าไฟฟ้า', 'Electricity', 3, 0, 1, '2026-05-16 09:43:29'),
(207, 1, 5, 205, '5309020000', 'ค่าน้ำประปา', 'Water Supply', 3, 0, 1, '2026-05-16 09:43:29'),
(208, 1, 5, 205, '5309030000', 'ค่าโทรศัพท์', 'Telephone', 3, 0, 1, '2026-05-16 09:43:29'),
(209, 1, 5, 205, '5309040000', 'ค่าอินเทอร์เน็ตและสื่อสาร', 'Internet and Communication', 3, 0, 1, '2026-05-16 09:43:29'),
(210, 1, 5, 205, '5309050000', 'ค่าวัสดุสำนักงาน', 'Office Supplies', 3, 0, 1, '2026-05-16 09:43:29'),
(211, 1, 5, 205, '5309060000', 'ค่าซ่อมแซมและบำรุงรักษา', 'Repairs and Maintenance', 3, 0, 1, '2026-05-16 09:43:29'),
(212, 1, 5, 205, '5309070000', 'ค่าน้ำมันเชื้อเพลิง', 'Fuel Expenses', 3, 0, 1, '2026-05-16 09:43:29'),
(213, 1, 5, 205, '5309080000', 'ค่าไปรษณีย์และค่าธรรมเนียมขนส่ง', 'Postage and Freight', 3, 0, 1, '2026-05-16 09:43:29'),
(214, 1, 5, 205, '5309090000', 'ค่าเบี้ยประกันภัย', 'Insurance Premiums', 3, 0, 1, '2026-05-16 09:43:29'),
(215, 1, 5, 193, '5310000000', 'ค่าธรรมเนียมและค่าบริการวิชาชีพ', 'Professional Fees and Service Charges', 2, 1, 1, '2026-05-16 09:43:29'),
(216, 1, 5, 215, '5310010000', 'ค่าธรรมเนียมธนาคาร', 'Bank Charges', 3, 0, 1, '2026-05-16 09:43:29'),
(217, 1, 5, 215, '5310020000', 'ค่าทำบัญชีและตรวจสอบบัญชี', 'Bookkeeping and Audit Fees', 3, 0, 1, '2026-05-16 09:43:29'),
(218, 1, 5, 215, '5310030000', 'ค่าที่ปรึกษากฎหมายและภาษี', 'Legal and Tax Consulting Fees', 3, 0, 1, '2026-05-16 09:43:29'),
(219, 1, 5, 215, '5310040000', 'ค่าธรรมเนียมราชการและภาษีอื่น', 'Government Fees and Other Taxes', 3, 0, 1, '2026-05-16 09:43:29'),
(220, 1, 5, 193, '5311000000', 'ค่าเสื่อมราคา', 'Depreciation and Amortization', 2, 1, 1, '2026-05-16 09:43:29'),
(221, 1, 5, 220, '5311010000', 'ค่าเสื่อมราคา-อาคาร', 'Depreciation - Buildings', 3, 0, 1, '2026-05-16 09:43:29'),
(222, 1, 5, 220, '5311020000', 'ค่าเสื่อมราคา-อุปกรณ์', 'Depreciation - Equipment', 3, 0, 1, '2026-05-16 09:43:29'),
(223, 1, 5, 220, '5311030000', 'ค่าตัดจำหน่าย-สินทรัพย์ไม่มีตัวตน', 'Amortization - Intangible Assets', 3, 0, 1, '2026-05-16 09:43:29'),
(224, 1, 5, 193, '5312000000', 'ส่วนขาดทุนจากประมาณการส่วนสูญเสียด้านเครดิตของลูกหนี้การค้า', 'Expected Credit Loss (ECL) - Trade Receivables', 2, 0, 1, '2026-05-16 09:43:29'),
(225, 1, 5, 193, '5313000000', 'ค่าใช้จ่ายเบ็ดเตล็ด', 'Miscellaneous Expenses', 2, 0, 1, '2026-05-16 09:43:29'),
(226, 1, 5, 193, '5314000000', 'ค่าใช้จ่ายอื่น', 'Other Expenses', 2, 1, 1, '2026-05-16 09:43:29'),
(227, 1, 5, 226, '5314010000', 'หนี้สูญ', 'Bad Debts', 3, 0, 1, '2026-05-16 09:43:29'),
(228, 1, 5, 226, '5314020000', 'เงินบริจาค', 'Donations', 3, 0, 1, '2026-05-16 09:43:29'),
(229, 1, 5, 226, '5314030000', 'ค่าอากรแสตมป์', 'Stamp Duty', 3, 0, 1, '2026-05-16 09:43:29'),
(230, 1, 5, 178, '5400000000', 'ต้นทุนทางการเงิน', 'Finance Costs', 1, 1, 1, '2026-05-16 09:43:29'),
(231, 1, 5, 230, '5401000000', 'ดอกเบี้ยจ่าย', 'Interest Expense', 2, 0, 1, '2026-05-16 09:43:29'),
(232, 1, 5, 230, '5402000000', 'ค่าธรรมเนียมเงินกู้', 'Loan Fees', 2, 0, 1, '2026-05-16 09:43:29'),
(233, 1, 5, 178, '5500000000', 'ภาษีเงินได้', 'Income Tax Expense', 1, 1, 1, '2026-05-16 09:43:29'),
(234, 1, 5, 233, '5501000000', 'ภาษีเงินได้นิติบุคคล', 'Corporate Income Tax', 2, 0, 1, '2026-05-16 09:43:29'),
(235, 1, 5, 233, '5502000000', 'ภาษีเงินได้รอตัดบัญชี', 'Deferred Income Tax', 2, 0, 1, '2026-05-16 09:43:29'),
(236, 1, 6, NULL, '6000000000', 'บัญชีพัก', 'Suspense / Clearing Accounts', 0, 1, 1, '2026-05-16 09:43:29'),
(237, 1, 6, 236, '6100000000', 'กำไร (ขาดทุน) จากการจำหน่ายสินทรัพย์', 'Gain (Loss) on Disposal of Assets', 1, 0, 1, '2026-05-16 09:43:29'),
(238, 1, 6, 236, '6200000000', 'กำไร (ขาดทุน) จากอัตราแลกเปลี่ยน', 'Gain (Loss) on Foreign Exchange', 1, 0, 1, '2026-05-16 09:43:29'),
(239, 1, 6, 236, '6300000000', 'ส่วนปรับปรุงเศษสตางค์จากการใช้เงินสด', 'Cash Rounding Adjustment', 1, 0, 1, '2026-05-16 09:43:29'),
(240, 1, 6, 236, '6400000000', 'เงินรับรอตรวจสอบ', 'Suspense - Receipts Pending Investigation', 1, 0, 1, '2026-05-16 09:43:29'),
(241, 1, 6, 236, '6500000000', 'เงินจ่ายรอตรวจสอบ', 'Suspense - Payments Pending Investigation', 1, 0, 1, '2026-05-16 09:43:29');

-- --------------------------------------------------------

--
-- Table structure for table `account_types`
--

CREATE TABLE `account_types` (
  `id` tinyint(3) UNSIGNED NOT NULL,
  `code` varchar(20) NOT NULL,
  `name_th` varchar(100) NOT NULL,
  `name_en` varchar(100) DEFAULT NULL,
  `normal_balance` enum('DR','CR') NOT NULL,
  `fs_section` enum('BS','PL','CF') NOT NULL COMMENT 'Balance Sheet / P&L / Cash Flow'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `account_types`
--

INSERT INTO `account_types` (`id`, `code`, `name_th`, `name_en`, `normal_balance`, `fs_section`) VALUES
(1, 'ASSET', 'สินทรัพย์', 'Asset', 'DR', 'BS'),
(2, 'LIABILITY', 'หนี้สิน', 'Liability', 'CR', 'BS'),
(3, 'EQUITY', 'ส่วนของผู้ถือหุ้น', 'Equity', 'CR', 'BS'),
(4, 'REVENUE', 'รายได้', 'Revenue', 'CR', 'PL'),
(5, 'EXPENSE', 'ค่าใช้จ่าย', 'Expense', 'DR', 'PL'),
(6, 'SUSPENSE', 'บัญชีพัก', 'Suspense', 'DR', 'PL');

-- --------------------------------------------------------

--
-- Table structure for table `approval_actions`
--

CREATE TABLE `approval_actions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `request_id` bigint(20) UNSIGNED NOT NULL,
  `step_no` tinyint(3) UNSIGNED NOT NULL,
  `actor_user_id` bigint(20) UNSIGNED NOT NULL COMMENT 'ผู้ดำเนินการ',
  `delegated_from` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'ถ้าเป็น delegate: user ที่มอบหมายให้',
  `action` enum('SUBMIT','APPROVE','REJECT','DELEGATE','RECALL','COMMENT','ESCALATE','AUTO_APPROVE') NOT NULL,
  `reason` text DEFAULT NULL COMMENT 'เหตุผลประกอบการ approve/reject',
  `digital_signature` text DEFAULT NULL COMMENT 'Digital signature (Base64) สร้างจาก private key ของ actor',
  `signature_key_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'FK → user_signatures.id (key ที่ใช้ sign)',
  `signed_payload` text DEFAULT NULL COMMENT 'payload ที่ถูก sign (canonical JSON string) เพื่อ verify ในอนาคต',
  `ip_address` varchar(45) DEFAULT NULL,
  `user_agent` varchar(500) DEFAULT NULL,
  `acted_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Immutable log ทุก action การอนุมัติ';

-- --------------------------------------------------------

--
-- Table structure for table `approval_delegates`
--

CREATE TABLE `approval_delegates` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `delegator_id` bigint(20) UNSIGNED NOT NULL COMMENT 'ผู้มอบหมาย',
  `delegate_id` bigint(20) UNSIGNED NOT NULL COMMENT 'ผู้รับมอบหมาย',
  `policy_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'NULL = มอบหมายทุก policy',
  `doc_module` varchar(30) DEFAULT NULL COMMENT 'NULL = ทุกประเภทเอกสาร',
  `valid_from` datetime NOT NULL,
  `valid_until` datetime NOT NULL,
  `reason` varchar(500) DEFAULT NULL COMMENT 'เหตุผล เช่น ลาพักร้อน 1-7 มี.ค.',
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='มอบหมายสิทธิ์อนุมัติชั่วคราว';

-- --------------------------------------------------------

--
-- Table structure for table `approval_notifications`
--

CREATE TABLE `approval_notifications` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `request_id` bigint(20) UNSIGNED NOT NULL,
  `recipient_id` bigint(20) UNSIGNED NOT NULL,
  `channel` enum('EMAIL','LINE','IN_APP','SMS') NOT NULL DEFAULT 'IN_APP',
  `noti_type` enum('APPROVAL_REQUESTED','APPROVED','REJECTED','SLA_WARNING','DELEGATED','RECALLED') NOT NULL,
  `sent_at` timestamp NULL DEFAULT NULL,
  `read_at` timestamp NULL DEFAULT NULL,
  `status` enum('PENDING','SENT','FAILED','READ') NOT NULL DEFAULT 'PENDING',
  `error_msg` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `approval_policies`
--

CREATE TABLE `approval_policies` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `doc_module` varchar(30) NOT NULL COMMENT 'PO=ใบสั่งซื้อ, SO=ใบสั่งขาย, INV=Invoice, PAY=Payment, GRN=รับสินค้า, JV=Journal, ADJ=ปรับสต็อก',
  `policy_name` varchar(255) NOT NULL,
  `description` text DEFAULT NULL,
  `conditions` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'JSON rules: amount_gte, amount_lte, department_ids, branch_ids, vendor_ids' CHECK (json_valid(`conditions`)),
  `total_steps` tinyint(3) UNSIGNED NOT NULL DEFAULT 1,
  `on_reject` enum('RETURN_TO_DRAFT','CANCEL','RETURN_TO_PREV_STEP') NOT NULL DEFAULT 'RETURN_TO_DRAFT',
  `sla_hours` smallint(5) UNSIGNED DEFAULT 48 COMMENT 'ชั่วโมงที่ต้องอนุมัติ ก่อนแจ้งเตือน',
  `sla_action` enum('NOTIFY_ONLY','AUTO_ESCALATE','AUTO_APPROVE') NOT NULL DEFAULT 'NOTIFY_ONLY',
  `priority` tinyint(3) UNSIGNED NOT NULL DEFAULT 10 COMMENT 'ลำดับ policy ถ้า doc ตรงหลาย policy ใช้ priority ต่ำสุดก่อน',
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Policy การอนุมัติต่อประเภทเอกสาร';

-- --------------------------------------------------------

--
-- Table structure for table `approval_requests`
--

CREATE TABLE `approval_requests` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `policy_id` bigint(20) UNSIGNED NOT NULL,
  `doc_module` varchar(30) NOT NULL COMMENT 'PO, SO, INV, PAY, GRN, JV, ADJ',
  `doc_id` bigint(20) UNSIGNED NOT NULL COMMENT 'ID ของเอกสารนั้น',
  `doc_no` varchar(50) NOT NULL COMMENT 'เลขที่เอกสาร (ไว้แสดง)',
  `doc_amount` decimal(18,2) DEFAULT NULL COMMENT 'มูลค่าเอกสาร (ใช้ match conditions)',
  `doc_date` date DEFAULT NULL,
  `status` enum('PENDING','APPROVED','REJECTED','CANCELLED','RECALLED') NOT NULL DEFAULT 'PENDING',
  `current_step` tinyint(3) UNSIGNED NOT NULL DEFAULT 1,
  `total_steps` tinyint(3) UNSIGNED NOT NULL,
  `requested_by` bigint(20) UNSIGNED NOT NULL,
  `requested_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `approved_at` timestamp NULL DEFAULT NULL,
  `rejected_at` timestamp NULL DEFAULT NULL,
  `sla_deadline` timestamp NULL DEFAULT NULL,
  `doc_snapshot` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'snapshot header ของเอกสาร ณ เวลา submit' CHECK (json_valid(`doc_snapshot`)),
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Instance การอนุมัติต่อ document จริง';

-- --------------------------------------------------------

--
-- Table structure for table `approval_steps`
--

CREATE TABLE `approval_steps` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `policy_id` bigint(20) UNSIGNED NOT NULL,
  `step_no` tinyint(3) UNSIGNED NOT NULL COMMENT '1, 2, 3...',
  `step_name` varchar(100) NOT NULL COMMENT 'เช่น หัวหน้าแผนก, CFO, MD',
  `approver_type` enum('SPECIFIC_USER','ROLE','DEPARTMENT_HEAD','DYNAMIC') NOT NULL DEFAULT 'ROLE',
  `approver_role_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'ใช้เมื่อ type = ROLE',
  `approver_user_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'ใช้เมื่อ type = SPECIFIC_USER',
  `approval_mode` enum('ANY_ONE','ALL_MUST') NOT NULL DEFAULT 'ANY_ONE',
  `can_skip` tinyint(1) NOT NULL DEFAULT 0,
  `allow_delegate` tinyint(1) NOT NULL DEFAULT 1,
  `require_reason_on_approve` tinyint(1) NOT NULL DEFAULT 0,
  `require_reason_on_reject` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='ขั้นตอนการอนุมัติภายใน policy';

-- --------------------------------------------------------

--
-- Table structure for table `audit_logs`
--

CREATE TABLE `audit_logs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `table_name` varchar(100) NOT NULL,
  `record_id` bigint(20) UNSIGNED NOT NULL,
  `action` enum('INSERT','UPDATE','DELETE','POST','APPROVE','CANCEL','REVERSE') NOT NULL,
  `old_values` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`old_values`)),
  `new_values` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`new_values`)),
  `ip_address` varchar(45) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Immutable audit trail - no updates/deletes allowed';

-- --------------------------------------------------------

--
-- Table structure for table `branches`
--

CREATE TABLE `branches` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `code` varchar(20) NOT NULL,
  `name_th` varchar(255) NOT NULL,
  `name_en` varchar(255) DEFAULT NULL,
  `branch_no` varchar(10) DEFAULT NULL COMMENT 'สาขาที่ for tax invoice',
  `address` text DEFAULT NULL,
  `phone` varchar(50) DEFAULT NULL,
  `is_hq` tinyint(1) NOT NULL DEFAULT 0,
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Multi-Branch';

-- --------------------------------------------------------

--
-- Table structure for table `business_partners`
--

CREATE TABLE `business_partners` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `code` varchar(30) NOT NULL,
  `name_th` varchar(255) NOT NULL,
  `name_en` varchar(255) DEFAULT NULL,
  `partner_type` set('CUSTOMER','VENDOR') NOT NULL,
  `tax_id` varchar(20) DEFAULT NULL,
  `branch_no` varchar(10) DEFAULT NULL COMMENT 'สาขา for tax invoice',
  `credit_days` smallint(6) DEFAULT 30,
  `credit_limit` decimal(18,2) DEFAULT 0.00,
  `billing_address` text DEFAULT NULL,
  `shipping_address` text DEFAULT NULL,
  `phone` varchar(100) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `ar_account_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'Default AR account',
  `ap_account_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'Default AP account',
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Unified customer & vendor master';

-- --------------------------------------------------------

--
-- Table structure for table `companies`
--

CREATE TABLE `companies` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `code` varchar(20) NOT NULL,
  `name_th` varchar(255) NOT NULL,
  `name_en` varchar(255) DEFAULT NULL,
  `tax_id` varchar(20) DEFAULT NULL,
  `registered_no` varchar(50) DEFAULT NULL,
  `address` text DEFAULT NULL,
  `phone` varchar(50) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `logo_url` varchar(500) DEFAULT NULL,
  `fiscal_year_start` tinyint(4) NOT NULL DEFAULT 1 COMMENT 'Month 1-12',
  `base_currency` char(3) NOT NULL DEFAULT 'THB',
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Multi-Company root entity';

--
-- Dumping data for table `companies`
--

INSERT INTO `companies` (`id`, `code`, `name_th`, `name_en`, `tax_id`, `registered_no`, `address`, `phone`, `email`, `logo_url`, `fiscal_year_start`, `base_currency`, `is_active`, `created_at`, `updated_at`) VALUES
(1, 'CPE', 'บริษัท เชิดชัย ไพร์ท เอ็นจิเนียริ่ง จำกัด', 'Choedchai Prime Engineering Company Limited', '0745569002816', '0745569002816', '72/7 หมู่ที่ 3 ตำบลท่าทราย อำเภอเมืองสมุทรสาคร จังหวัดสมุทรสาคร 74000', '', ' choedchai.en@gmail.com', '', 1, 'THB', 1, '2026-03-07 08:51:46', '2026-04-30 09:42:46');

-- --------------------------------------------------------

--
-- Table structure for table `currencies`
--

CREATE TABLE `currencies` (
  `code` char(3) NOT NULL,
  `name` varchar(100) NOT NULL,
  `symbol` varchar(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `currencies`
--

INSERT INTO `currencies` (`code`, `name`, `symbol`) VALUES
('CNY', 'Chinese Yuan', '¥'),
('EUR', 'Euro', '€'),
('JPY', 'Japanese Yen', '¥'),
('SGD', 'Singapore Dollar', 'S$'),
('THB', 'Thai Baht', '฿'),
('USD', 'US Dollar', '$');

-- --------------------------------------------------------

--
-- Table structure for table `delivery_orders`
--

CREATE TABLE `delivery_orders` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `do_no` varchar(30) NOT NULL,
  `so_id` bigint(20) UNSIGNED DEFAULT NULL,
  `customer_id` bigint(20) UNSIGNED NOT NULL,
  `delivery_date` date NOT NULL,
  `status` enum('DRAFT','SHIPPED','POSTED','CANCELLED') NOT NULL DEFAULT 'DRAFT',
  `journal_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `delivery_order_lines`
--

CREATE TABLE `delivery_order_lines` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `do_id` bigint(20) UNSIGNED NOT NULL,
  `so_line_id` bigint(20) UNSIGNED DEFAULT NULL,
  `product_id` bigint(20) UNSIGNED NOT NULL,
  `warehouse_id` bigint(20) UNSIGNED NOT NULL,
  `location_id` bigint(20) UNSIGNED DEFAULT NULL,
  `uom_id` bigint(20) UNSIGNED NOT NULL,
  `qty_shipped` decimal(18,4) NOT NULL,
  `unit_cost` decimal(18,4) NOT NULL DEFAULT 0.0000
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `departments`
--

CREATE TABLE `departments` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED DEFAULT NULL,
  `parent_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'Self-ref for dept hierarchy',
  `code` varchar(20) NOT NULL,
  `name_th` varchar(255) NOT NULL,
  `name_en` varchar(255) DEFAULT NULL,
  `cost_center` varchar(20) DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Multi-Department with hierarchy';

-- --------------------------------------------------------

--
-- Table structure for table `document_sequences`
--

CREATE TABLE `document_sequences` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED DEFAULT NULL,
  `doc_type` varchar(30) NOT NULL COMMENT 'SO, PO, INV, JV, GRN, DO, REC, PAY ...',
  `prefix` varchar(10) DEFAULT NULL,
  `year_format` enum('NONE','YY','YYYY') NOT NULL DEFAULT 'YY',
  `month_format` enum('NONE','MM') NOT NULL DEFAULT 'MM',
  `sep_char` varchar(5) DEFAULT '-',
  `padding` tinyint(4) NOT NULL DEFAULT 5,
  `current_number` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `reset_cycle` enum('NEVER','YEARLY','MONTHLY') NOT NULL DEFAULT 'YEARLY',
  `last_reset_date` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `document_stamps`
--

CREATE TABLE `document_stamps` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `request_id` bigint(20) UNSIGNED NOT NULL COMMENT 'FK → approval_requests',
  `doc_module` varchar(30) NOT NULL,
  `doc_id` bigint(20) UNSIGNED NOT NULL,
  `doc_no` varchar(50) NOT NULL,
  `stamp_uuid` char(36) NOT NULL COMMENT 'UUID พิมพ์เป็น QR บนเอกสาร',
  `content_hash` varchar(128) NOT NULL COMMENT 'SHA-256 hash ของ document content ณ เวลาพิมพ์',
  `hash_algorithm` varchar(20) NOT NULL DEFAULT 'SHA-256',
  `approvers_json` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'snapshot รายชื่อ+ตำแหน่ง+fingerprint ของผู้อนุมัติทุกขั้น' CHECK (json_valid(`approvers_json`)),
  `combined_signature` text DEFAULT NULL COMMENT 'Combined/chained digital signature จากทุก approver',
  `printed_by` bigint(20) UNSIGNED NOT NULL,
  `printed_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `print_count` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT 'นับจำนวนครั้งที่พิมพ์',
  `is_void` tinyint(1) NOT NULL DEFAULT 0,
  `voided_by` bigint(20) UNSIGNED DEFAULT NULL,
  `voided_at` timestamp NULL DEFAULT NULL,
  `void_reason` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Stamp ที่ embed ลงเอกสารพิมพ์ ใช้ verify ความถูกต้อง';

-- --------------------------------------------------------

--
-- Table structure for table `exchange_rates`
--

CREATE TABLE `exchange_rates` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `from_currency` char(3) NOT NULL,
  `to_currency` char(3) NOT NULL,
  `rate` decimal(18,6) NOT NULL,
  `rate_date` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `fiscal_periods`
--

CREATE TABLE `fiscal_periods` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `year` smallint(6) NOT NULL,
  `period` tinyint(4) NOT NULL COMMENT '1-12',
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `status` enum('OPEN','CLOSED','LOCKED') NOT NULL DEFAULT 'OPEN',
  `closed_by` bigint(20) UNSIGNED DEFAULT NULL,
  `closed_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `goods_receipts`
--

CREATE TABLE `goods_receipts` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `grn_no` varchar(30) NOT NULL,
  `po_id` bigint(20) UNSIGNED DEFAULT NULL,
  `vendor_id` bigint(20) UNSIGNED NOT NULL,
  `receipt_date` date NOT NULL,
  `status` enum('DRAFT','POSTED','CANCELLED') NOT NULL DEFAULT 'DRAFT',
  `journal_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `goods_receipt_lines`
--

CREATE TABLE `goods_receipt_lines` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `grn_id` bigint(20) UNSIGNED NOT NULL,
  `po_line_id` bigint(20) UNSIGNED DEFAULT NULL,
  `product_id` bigint(20) UNSIGNED NOT NULL,
  `warehouse_id` bigint(20) UNSIGNED NOT NULL,
  `location_id` bigint(20) UNSIGNED DEFAULT NULL,
  `uom_id` bigint(20) UNSIGNED NOT NULL,
  `qty_received` decimal(18,4) NOT NULL,
  `unit_cost` decimal(18,4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `inventory_balances`
--

CREATE TABLE `inventory_balances` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `warehouse_id` bigint(20) UNSIGNED NOT NULL,
  `location_id` bigint(20) UNSIGNED DEFAULT NULL,
  `product_id` bigint(20) UNSIGNED NOT NULL,
  `qty_on_hand` decimal(18,4) NOT NULL DEFAULT 0.0000,
  `qty_reserved` decimal(18,4) NOT NULL DEFAULT 0.0000,
  `qty_available` decimal(18,4) GENERATED ALWAYS AS (`qty_on_hand` - `qty_reserved`) VIRTUAL,
  `avg_cost` decimal(18,4) NOT NULL DEFAULT 0.0000,
  `last_updated` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `inventory_transactions`
--

CREATE TABLE `inventory_transactions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `warehouse_id` bigint(20) UNSIGNED NOT NULL,
  `location_id` bigint(20) UNSIGNED DEFAULT NULL,
  `product_id` bigint(20) UNSIGNED NOT NULL,
  `txn_type` enum('RECEIVE','ISSUE','ADJUST','TRANSFER_IN','TRANSFER_OUT','RETURN_IN','RETURN_OUT') NOT NULL,
  `ref_module` varchar(30) DEFAULT NULL COMMENT 'PO, SO, TRANSFER, ADJUST',
  `ref_id` bigint(20) UNSIGNED DEFAULT NULL,
  `ref_line_id` bigint(20) UNSIGNED DEFAULT NULL,
  `txn_date` date NOT NULL,
  `qty` decimal(18,4) NOT NULL,
  `unit_cost` decimal(18,4) NOT NULL DEFAULT 0.0000,
  `total_cost` decimal(18,4) GENERATED ALWAYS AS (`qty` * `unit_cost`) VIRTUAL,
  `journal_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'GL Journal created',
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `invoices`
--

CREATE TABLE `invoices` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `invoice_no` varchar(30) NOT NULL,
  `invoice_type` enum('BILLING_NOTE','TAX_INVOICE','RECEIPT_TAX_INVOICE','PURCHASE_INVOICE','DEBIT_NOTE','CREDIT_NOTE') NOT NULL DEFAULT 'BILLING_NOTE',
  `tax_invoice_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'FK to invoices.id ที่เป็น TAX_INVOICE (ถ้า billing note ถูกแปลงแล้ว)',
  `tax_invoice_date` date DEFAULT NULL COMMENT 'วันที่บนใบกำกับภาษี (ออกเมื่อรับชำระ)',
  `tax_invoice_no` varchar(30) DEFAULT NULL COMMENT 'เลขที่ใบกำกับภาษี (ถ้าแยกออกจาก billing note)',
  `payment_received_date` date DEFAULT NULL,
  `tax_on_payment` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'TRUE = ออกใบกำกับเมื่อรับชำระ (บริการ), FALSE = ออกพร้อมส่งของ (สินค้า)',
  `ref_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'SO or PO id',
  `ref_module` varchar(20) DEFAULT NULL COMMENT 'SO=Sales Order, PO=Purchase Order, BN=Billing Note, MANUAL',
  `partner_id` bigint(20) UNSIGNED NOT NULL,
  `invoice_date` date NOT NULL,
  `due_date` date NOT NULL,
  `currency_code` char(3) NOT NULL DEFAULT 'THB',
  `exchange_rate` decimal(18,6) NOT NULL DEFAULT 1.000000,
  `subtotal` decimal(18,2) NOT NULL DEFAULT 0.00,
  `discount_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `tax_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `total_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `paid_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `outstanding` decimal(18,2) GENERATED ALWAYS AS (`total_amount` - `paid_amount`) VIRTUAL,
  `status` enum('DRAFT','POSTED','PARTIAL','PAID','OVERDUE','CANCELLED','REVERSED') NOT NULL DEFAULT 'DRAFT',
  `journal_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `invoice_lines`
--

CREATE TABLE `invoice_lines` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `invoice_id` bigint(20) UNSIGNED NOT NULL,
  `line_no` smallint(6) NOT NULL,
  `product_id` bigint(20) UNSIGNED DEFAULT NULL,
  `description` varchar(500) NOT NULL,
  `uom_id` bigint(20) UNSIGNED DEFAULT NULL,
  `qty` decimal(18,4) NOT NULL DEFAULT 1.0000,
  `unit_price` decimal(18,4) NOT NULL,
  `discount_pct` decimal(5,2) NOT NULL DEFAULT 0.00,
  `line_amount` decimal(18,2) NOT NULL,
  `tax_type` enum('NO_TAX','VAT7','EXEMPT') NOT NULL DEFAULT 'VAT7',
  `tax_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `account_id` bigint(20) UNSIGNED DEFAULT NULL,
  `ref_line_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'SO line id หรือ PO line id ที่ดึงมา populate',
  `ref_module` varchar(20) DEFAULT NULL COMMENT 'SO / PO',
  `billed_qty` decimal(18,4) DEFAULT NULL COMMENT 'qty ที่แจ้งหนี้แล้ว (partial billing รองรับ)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `journals`
--

CREATE TABLE `journals` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `fiscal_period_id` bigint(20) UNSIGNED NOT NULL,
  `journal_no` varchar(30) NOT NULL,
  `journal_type` enum('MANUAL','SALES','PURCHASE','PAYMENT','RECEIPT','CLOSING','ADJUSTMENT') NOT NULL,
  `ref_module` varchar(30) DEFAULT NULL COMMENT 'Source module: SO, PO, AP, AR ...',
  `ref_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'Source document ID',
  `transaction_date` date NOT NULL,
  `description` varchar(500) DEFAULT NULL,
  `status` enum('DRAFT','POSTED','REVERSED') NOT NULL DEFAULT 'DRAFT',
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `posted_by` bigint(20) UNSIGNED DEFAULT NULL,
  `posted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `journal_lines`
--

CREATE TABLE `journal_lines` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `journal_id` bigint(20) UNSIGNED NOT NULL,
  `line_no` smallint(6) NOT NULL,
  `account_id` bigint(20) UNSIGNED NOT NULL,
  `department_id` bigint(20) UNSIGNED DEFAULT NULL,
  `debit` decimal(18,2) NOT NULL DEFAULT 0.00,
  `credit` decimal(18,2) NOT NULL DEFAULT 0.00,
  `description` varchar(500) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `payments`
--

CREATE TABLE `payments` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `payment_no` varchar(30) NOT NULL,
  `payment_type` enum('RECEIPT','PAYMENT') NOT NULL COMMENT 'RECEIPT=AR, PAYMENT=AP',
  `partner_id` bigint(20) UNSIGNED NOT NULL,
  `payment_date` date NOT NULL,
  `payment_method` enum('CASH','TRANSFER','CHEQUE','CREDIT_CARD','OTHER') NOT NULL DEFAULT 'TRANSFER',
  `bank_account` varchar(100) DEFAULT NULL,
  `cheque_no` varchar(50) DEFAULT NULL,
  `currency_code` char(3) NOT NULL DEFAULT 'THB',
  `exchange_rate` decimal(18,6) NOT NULL DEFAULT 1.000000,
  `amount` decimal(18,2) NOT NULL,
  `withholding_tax` decimal(18,2) NOT NULL DEFAULT 0.00,
  `wht_transaction_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'FK → wht_transactions.id (ถ้ามีการหัก WHT)',
  `wht_certificate_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'FK → wht_certificates.id (ใบ 50 ทวิ ที่ออกพร้อมการจ่าย)',
  `net_amount` decimal(18,2) NOT NULL,
  `status` enum('DRAFT','POSTED','CANCELLED') NOT NULL DEFAULT 'DRAFT',
  `journal_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `payment_allocations`
--

CREATE TABLE `payment_allocations` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `payment_id` bigint(20) UNSIGNED NOT NULL,
  `invoice_id` bigint(20) UNSIGNED NOT NULL,
  `allocated_amount` decimal(18,2) NOT NULL,
  `allocated_date` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `permissions`
--

CREATE TABLE `permissions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `module` varchar(50) NOT NULL COMMENT 'e.g. SALES, PURCHASE, INVENTORY, GL',
  `action` varchar(50) NOT NULL COMMENT 'e.g. VIEW, CREATE, EDIT, DELETE, APPROVE, POST',
  `description` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `permissions`
--

INSERT INTO `permissions` (`id`, `module`, `action`, `description`) VALUES
(1, 'COMPANY', 'VIEW', 'ดูข้อมูลบริษัท'),
(2, 'COMPANY', 'EDIT', 'แก้ไขข้อมูลบริษัท'),
(3, 'GL', 'VIEW', 'ดูสมุดรายวัน'),
(4, 'GL', 'CREATE', 'สร้างสมุดรายวัน'),
(5, 'GL', 'POST', 'Post สมุดรายวัน'),
(6, 'GL', 'REVERSE', 'กลับรายการ'),
(7, 'AR', 'VIEW', 'ดูลูกหนี้'),
(8, 'AR', 'CREATE', 'สร้างใบแจ้งหนี้ขาย'),
(9, 'AR', 'APPROVE', 'อนุมัติลูกหนี้'),
(10, 'AP', 'VIEW', 'ดูเจ้าหนี้'),
(11, 'AP', 'CREATE', 'สร้างใบแจ้งหนี้ซื้อ'),
(12, 'AP', 'APPROVE', 'อนุมัติเจ้าหนี้'),
(13, 'SALES', 'VIEW', 'ดูคำสั่งขาย'),
(14, 'SALES', 'CREATE', 'สร้างคำสั่งขาย'),
(15, 'SALES', 'APPROVE', 'อนุมัติคำสั่งขาย'),
(16, 'PURCHASE', 'VIEW', 'ดูใบสั่งซื้อ'),
(17, 'PURCHASE', 'CREATE', 'สร้างใบสั่งซื้อ'),
(18, 'PURCHASE', 'APPROVE', 'อนุมัติใบสั่งซื้อ'),
(19, 'INVENTORY', 'VIEW', 'ดูสต็อกสินค้า'),
(20, 'INVENTORY', 'ADJUST', 'ปรับยอดสต็อก'),
(21, 'REPORT', 'VIEW', 'ดูรายงาน'),
(22, 'REPORT', 'EXPORT', 'Export รายงาน'),
(23, 'ADMIN', 'ALL', 'สิทธิ์ผู้ดูแลระบบทั้งหมด'),
(24, 'APPROVAL', 'VIEW', 'ดูรายการรออนุมัติ'),
(25, 'APPROVAL', 'SUBMIT', 'ส่งเอกสารขออนุมัติ'),
(26, 'APPROVAL', 'APPROVE', 'อนุมัติเอกสาร'),
(27, 'APPROVAL', 'REJECT', 'ปฏิเสธเอกสาร'),
(28, 'APPROVAL', 'DELEGATE', 'มอบหมายสิทธิ์อนุมัติ'),
(29, 'APPROVAL', 'RECALL', 'ดึงเอกสารกลับ'),
(30, 'APPROVAL', 'MANAGE_POLICY', 'จัดการ Policy การอนุมัติ'),
(31, 'SIGNATURE', 'MANAGE', 'จัดการลายเซ็นและ Digital Key'),
(32, 'DOCUMENT', 'PRINT', 'พิมพ์เอกสาร'),
(33, 'DOCUMENT', 'VERIFY_STAMP', 'ตรวจสอบความถูกต้องของเอกสาร');

-- --------------------------------------------------------

--
-- Table structure for table `po_billing_tracking`
--

CREATE TABLE `po_billing_tracking` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `po_id` bigint(20) UNSIGNED NOT NULL,
  `po_line_id` bigint(20) UNSIGNED NOT NULL,
  `invoice_id` bigint(20) UNSIGNED NOT NULL,
  `invoice_line_id` bigint(20) UNSIGNED NOT NULL,
  `billed_qty` decimal(18,4) NOT NULL,
  `billed_amount` decimal(18,2) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Track qty ที่รับใบแจ้งหนี้แล้วต่อ PO line';

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `category_id` bigint(20) UNSIGNED DEFAULT NULL,
  `code` varchar(50) NOT NULL,
  `name_th` varchar(255) NOT NULL,
  `name_en` varchar(255) DEFAULT NULL,
  `product_type` enum('PRODUCT','SERVICE','RAW_MATERIAL','FINISHED_GOOD') NOT NULL DEFAULT 'PRODUCT',
  `base_uom_id` bigint(20) UNSIGNED NOT NULL,
  `standard_cost` decimal(18,4) DEFAULT 0.0000,
  `list_price` decimal(18,4) DEFAULT 0.0000,
  `tax_type` enum('NO_TAX','VAT7','EXEMPT') NOT NULL DEFAULT 'VAT7',
  `is_inventory` tinyint(1) NOT NULL DEFAULT 1,
  `inventory_account_id` bigint(20) UNSIGNED DEFAULT NULL,
  `cogs_account_id` bigint(20) UNSIGNED DEFAULT NULL,
  `sales_account_id` bigint(20) UNSIGNED DEFAULT NULL,
  `purchase_account_id` bigint(20) UNSIGNED DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `product_categories`
--

CREATE TABLE `product_categories` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `parent_id` bigint(20) UNSIGNED DEFAULT NULL,
  `code` varchar(30) NOT NULL,
  `name_th` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `purchase_orders`
--

CREATE TABLE `purchase_orders` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `department_id` bigint(20) UNSIGNED DEFAULT NULL,
  `po_no` varchar(30) NOT NULL,
  `vendor_id` bigint(20) UNSIGNED NOT NULL,
  `po_date` date NOT NULL,
  `expected_date` date DEFAULT NULL,
  `currency_code` char(3) NOT NULL DEFAULT 'THB',
  `exchange_rate` decimal(18,6) NOT NULL DEFAULT 1.000000,
  `status` enum('DRAFT','CONFIRMED','PARTIAL','RECEIVED','CANCELLED') NOT NULL DEFAULT 'DRAFT',
  `subtotal` decimal(18,2) NOT NULL DEFAULT 0.00,
  `tax_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `total_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `notes` text DEFAULT NULL,
  `approved_by` bigint(20) UNSIGNED DEFAULT NULL,
  `approved_at` timestamp NULL DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `purchase_order_lines`
--

CREATE TABLE `purchase_order_lines` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `po_id` bigint(20) UNSIGNED NOT NULL,
  `line_no` smallint(6) NOT NULL,
  `product_id` bigint(20) UNSIGNED NOT NULL,
  `description` varchar(500) DEFAULT NULL,
  `uom_id` bigint(20) UNSIGNED NOT NULL,
  `qty_ordered` decimal(18,4) NOT NULL,
  `qty_received` decimal(18,4) NOT NULL DEFAULT 0.0000,
  `unit_price` decimal(18,4) NOT NULL,
  `discount_pct` decimal(5,2) NOT NULL DEFAULT 0.00,
  `line_amount` decimal(18,2) NOT NULL,
  `tax_type` enum('NO_TAX','VAT7','EXEMPT') NOT NULL DEFAULT 'VAT7',
  `tax_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `warehouse_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` text DEFAULT NULL,
  `is_system` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'System roles cannot be deleted',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `company_id`, `code`, `name`, `description`, `is_system`, `created_at`) VALUES
(1, 1, 'AC', 'Account Officer', NULL, 0, '2026-04-30 09:38:35');

-- --------------------------------------------------------

--
-- Table structure for table `role_permissions`
--

CREATE TABLE `role_permissions` (
  `role_id` bigint(20) UNSIGNED NOT NULL,
  `permission_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `role_permissions`
--

INSERT INTO `role_permissions` (`role_id`, `permission_id`) VALUES
(1, 23);

-- --------------------------------------------------------

--
-- Table structure for table `sales_orders`
--

CREATE TABLE `sales_orders` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `department_id` bigint(20) UNSIGNED DEFAULT NULL,
  `so_no` varchar(30) NOT NULL,
  `customer_id` bigint(20) UNSIGNED NOT NULL,
  `order_date` date NOT NULL,
  `delivery_date` date DEFAULT NULL,
  `currency_code` char(3) NOT NULL DEFAULT 'THB',
  `exchange_rate` decimal(18,6) NOT NULL DEFAULT 1.000000,
  `status` enum('DRAFT','CONFIRMED','PARTIAL','SHIPPED','INVOICED','CANCELLED') NOT NULL DEFAULT 'DRAFT',
  `subtotal` decimal(18,2) NOT NULL DEFAULT 0.00,
  `discount_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `tax_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `total_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `notes` text DEFAULT NULL,
  `salesperson_id` bigint(20) UNSIGNED DEFAULT NULL,
  `approved_by` bigint(20) UNSIGNED DEFAULT NULL,
  `approved_at` timestamp NULL DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sales_order_lines`
--

CREATE TABLE `sales_order_lines` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `so_id` bigint(20) UNSIGNED NOT NULL,
  `line_no` smallint(6) NOT NULL,
  `product_id` bigint(20) UNSIGNED NOT NULL,
  `description` varchar(500) DEFAULT NULL,
  `uom_id` bigint(20) UNSIGNED NOT NULL,
  `qty_ordered` decimal(18,4) NOT NULL,
  `qty_shipped` decimal(18,4) NOT NULL DEFAULT 0.0000,
  `unit_price` decimal(18,4) NOT NULL,
  `discount_pct` decimal(5,2) NOT NULL DEFAULT 0.00,
  `line_amount` decimal(18,2) NOT NULL,
  `tax_type` enum('NO_TAX','VAT7','EXEMPT') NOT NULL DEFAULT 'VAT7',
  `tax_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `warehouse_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `so_billing_tracking`
--

CREATE TABLE `so_billing_tracking` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `so_id` bigint(20) UNSIGNED NOT NULL,
  `so_line_id` bigint(20) UNSIGNED NOT NULL,
  `invoice_id` bigint(20) UNSIGNED NOT NULL,
  `invoice_line_id` bigint(20) UNSIGNED NOT NULL,
  `billed_qty` decimal(18,4) NOT NULL,
  `billed_amount` decimal(18,2) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Track qty ที่แจ้งหนี้แล้วต่อ SO line (partial billing)';

-- --------------------------------------------------------

--
-- Table structure for table `tax_invoice_queue`
--

CREATE TABLE `tax_invoice_queue` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `billing_note_id` bigint(20) UNSIGNED NOT NULL COMMENT 'Invoice ที่เป็น BILLING_NOTE',
  `payment_id` bigint(20) UNSIGNED NOT NULL COMMENT 'Payment ที่ trigger การออก tax invoice',
  `status` enum('PENDING','PROCESSING','DONE','FAILED') NOT NULL DEFAULT 'PENDING',
  `tax_invoice_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'Tax Invoice ที่สร้างแล้ว',
  `error_msg` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `processed_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Queue สำหรับสร้าง Tax Invoice อัตโนมัติเมื่อรับชำระ (บริการ)';

-- --------------------------------------------------------

--
-- Table structure for table `units_of_measure`
--

CREATE TABLE `units_of_measure` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `code` varchar(20) NOT NULL,
  `name_th` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `units_of_measure`
--

INSERT INTO `units_of_measure` (`id`, `company_id`, `code`, `name_th`) VALUES
(1, 1, '01-01', 'Job');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `username` varchar(100) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `employee_code` varchar(50) DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `last_login_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password_hash`, `full_name`, `employee_code`, `is_active`, `last_login_at`, `created_at`, `updated_at`) VALUES
(2, 'bxston', 'choedchai.en@gmail.com', '$2a$10$dK/8g.4D7rdLkgcKmmgY6eq4IRrp9PDYpmMl0LXo6kelfkls982Iq', 'Prasit Krongmaroeng', '', 1, NULL, '2026-05-11 09:31:46', '2026-05-11 09:31:46');

-- --------------------------------------------------------

--
-- Table structure for table `user_company_access`
--

CREATE TABLE `user_company_access` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'NULL = all branches',
  `department_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'NULL = all departments',
  `role_id` bigint(20) UNSIGNED NOT NULL,
  `is_default` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Data Isolation: users can only see data within their access scope';

--
-- Dumping data for table `user_company_access`
--

INSERT INTO `user_company_access` (`id`, `user_id`, `company_id`, `branch_id`, `department_id`, `role_id`, `is_default`) VALUES
(1, 2, 1, NULL, NULL, 1, 0);

-- --------------------------------------------------------

--
-- Table structure for table `user_signatures`
--

CREATE TABLE `user_signatures` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `signature_image_url` varchar(1000) DEFAULT NULL COMMENT 'URL ลายเซ็นภาพ (PNG transparent background)',
  `signature_image_b64` mediumtext DEFAULT NULL COMMENT 'Base64 ลายเซ็น (ใช้กรณีไม่มี object storage)',
  `display_name_th` varchar(255) DEFAULT NULL COMMENT 'ชื่อที่แสดงบนเอกสาร เช่น นายสมชาย ใจดี',
  `display_name_en` varchar(255) DEFAULT NULL,
  `position_title_th` varchar(255) DEFAULT NULL COMMENT 'ตำแหน่งที่แสดง เช่น ผู้จัดการฝ่ายบัญชี',
  `position_title_en` varchar(255) DEFAULT NULL,
  `public_key` text DEFAULT NULL COMMENT 'RSA/ECDSA Public Key (PEM format) สำหรับ verify signature',
  `key_fingerprint` varchar(128) DEFAULT NULL COMMENT 'SHA-256 fingerprint ของ public key (แสดงบนเอกสาร)',
  `key_algorithm` enum('RSA_2048','RSA_4096','ECDSA_P256','ECDSA_P384') NOT NULL DEFAULT 'RSA_2048',
  `key_issued_at` timestamp NULL DEFAULT NULL,
  `key_expires_at` timestamp NULL DEFAULT NULL COMMENT 'วันหมดอายุ key (ควร rotate ทุก 1-2 ปี)',
  `key_revoked_at` timestamp NULL DEFAULT NULL COMMENT 'ถ้าถูก revoke จะไม่สามารถอนุมัติเพิ่มได้',
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='ลายเซ็นและ Digital Key ของผู้ใช้';

-- --------------------------------------------------------

--
-- Table structure for table `vat_rates`
--

CREATE TABLE `vat_rates` (
  `id` tinyint(3) UNSIGNED NOT NULL,
  `code` varchar(20) NOT NULL COMMENT 'VAT7, VAT0, EXEMPT, OUT_OF_SCOPE',
  `name_th` varchar(100) NOT NULL,
  `name_en` varchar(100) DEFAULT NULL,
  `rate` decimal(5,2) NOT NULL COMMENT 'เปอร์เซ็นต์ เช่น 7.00, 0.00',
  `vat_category` enum('STANDARD','ZERO_RATED','EXEMPT','OUT_OF_SCOPE') NOT NULL,
  `effective_from` date NOT NULL,
  `effective_until` date DEFAULT NULL COMMENT 'NULL = ยังใช้อยู่',
  `output_vat_account_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'ลูกหนี้ภาษีขาย',
  `input_vat_account_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'ลูกหนี้ภาษีซื้อ',
  `is_active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='อัตรา VAT และ GL mapping';

--
-- Dumping data for table `vat_rates`
--

INSERT INTO `vat_rates` (`id`, `code`, `name_th`, `name_en`, `rate`, `vat_category`, `effective_from`, `effective_until`, `output_vat_account_id`, `input_vat_account_id`, `is_active`) VALUES
(1, 'VAT7', 'ภาษีมูลค่าเพิ่ม 7%', 'Standard VAT 7%', 7.00, 'STANDARD', '1992-01-01', NULL, NULL, NULL, 1),
(2, 'VAT0', 'ภาษีมูลค่าเพิ่ม 0%', 'Zero-Rated VAT', 0.00, 'ZERO_RATED', '1992-01-01', NULL, NULL, NULL, 1),
(3, 'EXEMPT', 'ยกเว้นภาษีมูลค่าเพิ่ม', 'VAT Exempt', 0.00, 'EXEMPT', '1992-01-01', NULL, NULL, NULL, 1),
(4, 'NOSCOPE', 'นอกขอบเขต VAT', 'Out of Scope', 0.00, 'OUT_OF_SCOPE', '1992-01-01', NULL, NULL, NULL, 1);

-- --------------------------------------------------------

--
-- Table structure for table `vat_report_periods`
--

CREATE TABLE `vat_report_periods` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL COMMENT 'สาขาที่ยื่น ภ.พ.30',
  `report_year` smallint(6) NOT NULL,
  `report_month` tinyint(4) NOT NULL COMMENT '1-12',
  `due_date` date NOT NULL COMMENT 'วันครบกำหนดยื่น (ปกติ ภายใน 15 ของเดือนถัดไป)',
  `total_output_base` decimal(18,2) NOT NULL DEFAULT 0.00 COMMENT 'ยอดขาย (ฐาน)',
  `total_output_vat` decimal(18,2) NOT NULL DEFAULT 0.00 COMMENT 'ภาษีขาย',
  `total_input_base` decimal(18,2) NOT NULL DEFAULT 0.00 COMMENT 'ยอดซื้อ (ฐาน)',
  `total_input_vat` decimal(18,2) NOT NULL DEFAULT 0.00 COMMENT 'ภาษีซื้อ',
  `vat_payable` decimal(18,2) GENERATED ALWAYS AS (`total_output_vat` - `total_input_vat`) VIRTUAL COMMENT 'VAT ที่ต้องชำระ (+ = จ่าย, – = ขอคืน)',
  `status` enum('DRAFT','SUBMITTED','AMENDED') NOT NULL DEFAULT 'DRAFT',
  `submitted_at` timestamp NULL DEFAULT NULL,
  `submitted_by` bigint(20) UNSIGNED DEFAULT NULL,
  `reference_no` varchar(50) DEFAULT NULL COMMENT 'เลขที่อ้างอิงจากกรมสรรพากร',
  `payment_date` date DEFAULT NULL,
  `payment_amount` decimal(18,2) DEFAULT NULL,
  `penalty_amount` decimal(18,2) DEFAULT 0.00 COMMENT 'ค่าปรับ (ถ้ายื่นล่าช้า)',
  `surcharge_amount` decimal(18,2) DEFAULT 0.00 COMMENT 'เงินเพิ่ม',
  `notes` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='รายงาน ภ.พ.30 รายเดือนต่อสาขา';

-- --------------------------------------------------------

--
-- Stand-in structure for view `v_approval_history`
-- (See below for the actual view)
--
CREATE TABLE `v_approval_history` (
`doc_module` varchar(30)
,`doc_id` bigint(20) unsigned
,`doc_no` varchar(50)
,`step_no` tinyint(3) unsigned
,`step_name` varchar(100)
,`action` enum('SUBMIT','APPROVE','REJECT','DELEGATE','RECALL','COMMENT','ESCALATE','AUTO_APPROVE')
,`actor_name` varchar(255)
,`actor_display_name` varchar(255)
,`actor_position` varchar(255)
,`key_fingerprint` varchar(128)
,`reason` text
,`acted_at` timestamp
,`has_digital_signature` int(1)
);

-- --------------------------------------------------------

--
-- Stand-in structure for view `v_approval_inbox`
-- (See below for the actual view)
--
CREATE TABLE `v_approval_inbox` (
`request_id` bigint(20) unsigned
,`company_id` bigint(20) unsigned
,`doc_module` varchar(30)
,`doc_no` varchar(50)
,`doc_amount` decimal(18,2)
,`doc_date` date
,`current_step` tinyint(3) unsigned
,`total_steps` tinyint(3) unsigned
,`status` enum('PENDING','APPROVED','REJECTED','CANCELLED','RECALLED')
,`sla_deadline` timestamp
,`requested_at` timestamp
,`requested_by_name` varchar(255)
,`current_step_name` varchar(100)
,`approver_type` enum('SPECIFIC_USER','ROLE','DEPARTMENT_HEAD','DYNAMIC')
,`approver_role_id` bigint(20) unsigned
,`approver_user_id` bigint(20) unsigned
,`is_overdue` int(1)
);

-- --------------------------------------------------------

--
-- Table structure for table `warehouses`
--

CREATE TABLE `warehouses` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `code` varchar(20) NOT NULL,
  `name_th` varchar(255) NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `warehouse_locations`
--

CREATE TABLE `warehouse_locations` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `warehouse_id` bigint(20) UNSIGNED NOT NULL,
  `code` varchar(30) NOT NULL,
  `name` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `wht_certificates`
--

CREATE TABLE `wht_certificates` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `cert_no` varchar(30) NOT NULL COMMENT 'เลขที่ใบรับรองการหักภาษี ณ ที่จ่าย',
  `cert_date` date NOT NULL COMMENT 'วันที่ออกใบรับรอง',
  `partner_id` bigint(20) UNSIGNED NOT NULL,
  `partner_name` varchar(255) NOT NULL,
  `partner_tax_id` varchar(20) NOT NULL,
  `partner_address` text DEFAULT NULL,
  `payer_name` varchar(255) NOT NULL COMMENT 'ชื่อผู้จ่ายเงิน',
  `payer_tax_id` varchar(20) NOT NULL,
  `payer_address` text DEFAULT NULL,
  `total_income` decimal(18,2) NOT NULL DEFAULT 0.00,
  `total_wht` decimal(18,2) NOT NULL DEFAULT 0.00,
  `condition_type` enum('ONE_TIME','EVERY_TIME','OTHERS') NOT NULL DEFAULT 'EVERY_TIME',
  `status` enum('DRAFT','ISSUED','CANCELLED') NOT NULL DEFAULT 'DRAFT',
  `issued_by` bigint(20) UNSIGNED DEFAULT NULL,
  `issued_at` timestamp NULL DEFAULT NULL,
  `e_wht_ref_no` varchar(100) DEFAULT NULL COMMENT 'เลขที่อ้างอิงจากระบบ e-WHT ของกรมสรรพากร',
  `stamp_uuid` char(36) DEFAULT NULL,
  `content_hash` varchar(128) DEFAULT NULL,
  `notes` text DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='ใบรับรองการหักภาษี ณ ที่จ่าย (ใบ 50 ทวิ)';

-- --------------------------------------------------------

--
-- Table structure for table `wht_report_periods`
--

CREATE TABLE `wht_report_periods` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `form_type` enum('ND1','ND3','ND53') NOT NULL,
  `report_year` smallint(6) NOT NULL,
  `report_month` tinyint(4) NOT NULL,
  `due_date` date NOT NULL COMMENT 'ครบกำหนดยื่น (ปกติวันที่ 7 ของเดือนถัดไป, e-filing วันที่ 15)',
  `total_income` decimal(18,2) NOT NULL DEFAULT 0.00,
  `total_wht` decimal(18,2) NOT NULL DEFAULT 0.00,
  `penalty_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `surcharge_amount` decimal(18,2) NOT NULL DEFAULT 0.00,
  `status` enum('DRAFT','SUBMITTED','AMENDED') NOT NULL DEFAULT 'DRAFT',
  `submitted_at` timestamp NULL DEFAULT NULL,
  `submitted_by` bigint(20) UNSIGNED DEFAULT NULL,
  `reference_no` varchar(50) DEFAULT NULL COMMENT 'เลขที่อ้างอิงจากกรมสรรพากร',
  `payment_date` date DEFAULT NULL,
  `payment_amount` decimal(18,2) DEFAULT NULL,
  `notes` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='รายงาน ภ.ง.ด. รายเดือน (1/3/53)';

-- --------------------------------------------------------

--
-- Table structure for table `wht_transactions`
--

CREATE TABLE `wht_transactions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `company_id` bigint(20) UNSIGNED NOT NULL,
  `branch_id` bigint(20) UNSIGNED NOT NULL,
  `fiscal_period_id` bigint(20) UNSIGNED NOT NULL,
  `wht_direction` enum('WITHHELD','SUFFERED') DEFAULT NULL COMMENT 'WITHHELD=เราหักจากคนอื่น (จ่าย), SUFFERED=ถูกหักจากเรา (รับ)',
  `ref_module` varchar(30) NOT NULL COMMENT 'PAY=Payment, INV=Invoice',
  `ref_id` bigint(20) UNSIGNED NOT NULL,
  `ref_doc_no` varchar(50) NOT NULL,
  `payment_date` date NOT NULL COMMENT 'วันที่จ่ายเงิน (ใช้ระบุเดือนที่นำส่ง)',
  `partner_id` bigint(20) UNSIGNED NOT NULL,
  `partner_name` varchar(255) NOT NULL COMMENT 'snapshot',
  `partner_tax_id` varchar(20) NOT NULL COMMENT 'snapshot เลขผู้เสียภาษี',
  `partner_type` enum('INDIVIDUAL','JURISTIC') NOT NULL COMMENT 'บุคคลธรรมดา/นิติบุคคล → กำหนดแบบ ภ.ง.ด.',
  `wht_type_id` tinyint(3) UNSIGNED NOT NULL,
  `wht_description` varchar(255) DEFAULT NULL COMMENT 'รายละเอียดที่จะปรากฏในใบ 50 ทวิ',
  `income_amount` decimal(18,2) NOT NULL COMMENT 'เงินได้พึงประเมิน (ก่อนหัก)',
  `wht_rate` decimal(5,2) NOT NULL COMMENT 'อัตรา ณ เวลานั้น (snapshot)',
  `wht_amount` decimal(18,2) NOT NULL COMMENT 'จำนวนภาษีที่หัก',
  `net_paid_amount` decimal(18,2) GENERATED ALWAYS AS (`income_amount` - `wht_amount`) VIRTUAL COMMENT 'จำนวนเงินที่จ่ายจริง (หลังหัก)',
  `currency_code` char(3) NOT NULL DEFAULT 'THB',
  `exchange_rate` decimal(18,6) NOT NULL DEFAULT 1.000000,
  `income_amount_thb` decimal(18,2) NOT NULL,
  `wht_amount_thb` decimal(18,2) NOT NULL,
  `journal_id` bigint(20) UNSIGNED DEFAULT NULL,
  `wht_account_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'ภาษีหัก ณ ที่จ่ายค้างชำระ',
  `certificate_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'FK → wht_certificates.id',
  `is_remitted` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'นำส่งกรมสรรพากรแล้วหรือยัง',
  `report_period_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'FK → wht_report_periods.id',
  `created_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='บันทึก WHT ทุกรายการ (หัก ณ ที่จ่าย)';

-- --------------------------------------------------------

--
-- Table structure for table `wht_types`
--

CREATE TABLE `wht_types` (
  `id` tinyint(3) UNSIGNED NOT NULL,
  `code` varchar(20) NOT NULL,
  `section` varchar(20) NOT NULL COMMENT 'เช่น 40(1), 40(2), 40(3), 40(4)ก, 40(4)ข, 40(5), 40(6), 40(8)',
  `description_th` varchar(255) NOT NULL,
  `description_en` varchar(255) DEFAULT NULL,
  `default_rate` decimal(5,2) NOT NULL COMMENT 'อัตราปกติ %',
  `reduced_rate` decimal(5,2) DEFAULT NULL COMMENT 'อัตราลดหย่อน % (เช่น ตาม BOI, COVID มาตรการ)',
  `form_type` enum('ND1','ND3','ND53') DEFAULT NULL COMMENT 'ภ.ง.ด.1=เงินเดือน, ภ.ง.ด.3=บุคคลธรรมดา, ภ.ง.ด.53=นิติบุคคล',
  `is_active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='ประเภทเงินได้ตาม พรบ. ภาษีเงินได้';

--
-- Dumping data for table `wht_types`
--

INSERT INTO `wht_types` (`id`, `code`, `section`, `description_th`, `description_en`, `default_rate`, `reduced_rate`, `form_type`, `is_active`) VALUES
(1, '40_1', '40(1)', 'เงินเดือน ค่าจ้าง บำนาญ', 'Salary, Wages, Pension', 0.00, NULL, 'ND1', 1),
(2, '40_2', '40(2)', 'ค่าจ้างทำของ ค่าบริการ (บุคคล)', 'Service fee (individual)', 3.00, NULL, 'ND3', 1),
(3, '40_3', '40(3)', 'ค่าลิขสิทธิ์ สิทธิบัตร', 'Royalties, Patents', 3.00, NULL, 'ND3', 1),
(4, '40_4A', '40(4)ก', 'ดอกเบี้ยเงินฝาก พันธบัตร', 'Interest (deposits, bonds)', 15.00, NULL, 'ND3', 1),
(5, '40_4B', '40(4)ข', 'เงินปันผล', 'Dividends', 10.00, NULL, 'ND3', 1),
(6, '40_5', '40(5)', 'ค่าเช่า (บุคคล)', 'Rental income (individual)', 5.00, NULL, 'ND3', 1),
(7, '40_6', '40(6)', 'วิชาชีพอิสระ (แพทย์ กฎหมาย บัญชี วิศวกร)', 'Professional services (individual)', 3.00, NULL, 'ND3', 1),
(8, '40_7', '40(7)', 'รับเหมาแรงงาน', 'Contract labor', 3.00, NULL, 'ND3', 1),
(9, '40_8_SVC', '40(8)', 'ค่าบริการ/ค่าจ้าง (นิติบุคคล)', 'Service fee (juristic)', 3.00, NULL, 'ND53', 1),
(10, '40_8_ADV', '40(8)', 'ค่าโฆษณา', 'Advertising fee', 2.00, NULL, 'ND53', 1),
(11, '40_8_TRN', '40(8)', 'ค่าขนส่ง', 'Transportation fee', 1.00, NULL, 'ND53', 1),
(12, '40_8_INS', '40(8)', 'ค่าเบี้ยประกันภัย', 'Insurance premium', 1.00, NULL, 'ND53', 1),
(13, '40_8_RNT', '40(8)', 'ค่าเช่า (นิติบุคคล)', 'Rental income (juristic)', 5.00, NULL, 'ND53', 1);

-- --------------------------------------------------------

--
-- Structure for view `v_approval_history`
--
DROP TABLE IF EXISTS `v_approval_history`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `v_approval_history`  AS SELECT `ar`.`doc_module` AS `doc_module`, `ar`.`doc_id` AS `doc_id`, `ar`.`doc_no` AS `doc_no`, `aa`.`step_no` AS `step_no`, `aps`.`step_name` AS `step_name`, `aa`.`action` AS `action`, `u`.`full_name` AS `actor_name`, `us`.`display_name_th` AS `actor_display_name`, `us`.`position_title_th` AS `actor_position`, `us`.`key_fingerprint` AS `key_fingerprint`, `aa`.`reason` AS `reason`, `aa`.`acted_at` AS `acted_at`, `aa`.`digital_signature` is not null AS `has_digital_signature` FROM ((((`approval_actions` `aa` join `approval_requests` `ar` on(`ar`.`id` = `aa`.`request_id`)) join `approval_steps` `aps` on(`aps`.`policy_id` = `ar`.`policy_id` and `aps`.`step_no` = `aa`.`step_no`)) join `users` `u` on(`u`.`id` = `aa`.`actor_user_id`)) left join `user_signatures` `us` on(`us`.`user_id` = `aa`.`actor_user_id`)) ORDER BY `aa`.`acted_at` ASC ;

-- --------------------------------------------------------

--
-- Structure for view `v_approval_inbox`
--
DROP TABLE IF EXISTS `v_approval_inbox`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `v_approval_inbox`  AS SELECT `ar`.`id` AS `request_id`, `ar`.`company_id` AS `company_id`, `ar`.`doc_module` AS `doc_module`, `ar`.`doc_no` AS `doc_no`, `ar`.`doc_amount` AS `doc_amount`, `ar`.`doc_date` AS `doc_date`, `ar`.`current_step` AS `current_step`, `ar`.`total_steps` AS `total_steps`, `ar`.`status` AS `status`, `ar`.`sla_deadline` AS `sla_deadline`, `ar`.`requested_at` AS `requested_at`, `u`.`full_name` AS `requested_by_name`, `aps`.`step_name` AS `current_step_name`, `aps`.`approver_type` AS `approver_type`, `aps`.`approver_role_id` AS `approver_role_id`, `aps`.`approver_user_id` AS `approver_user_id`, CASE WHEN `ar`.`sla_deadline` < current_timestamp() THEN 1 ELSE 0 END AS `is_overdue` FROM (((`approval_requests` `ar` join `approval_policies` `apo` on(`apo`.`id` = `ar`.`policy_id`)) join `approval_steps` `aps` on(`aps`.`policy_id` = `ar`.`policy_id` and `aps`.`step_no` = `ar`.`current_step`)) join `users` `u` on(`u`.`id` = `ar`.`requested_by`)) WHERE `ar`.`status` = 'PENDING' ;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `accounts`
--
ALTER TABLE `accounts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_account_code` (`company_id`,`code`),
  ADD KEY `account_type_id` (`account_type_id`),
  ADD KEY `parent_id` (`parent_id`);

--
-- Indexes for table `account_types`
--
ALTER TABLE `account_types`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Indexes for table `approval_actions`
--
ALTER TABLE `approval_actions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `request_id` (`request_id`),
  ADD KEY `actor_user_id` (`actor_user_id`),
  ADD KEY `delegated_from` (`delegated_from`),
  ADD KEY `signature_key_id` (`signature_key_id`);

--
-- Indexes for table `approval_delegates`
--
ALTER TABLE `approval_delegates`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_delegate_active` (`company_id`,`delegate_id`,`valid_from`,`valid_until`),
  ADD KEY `delegator_id` (`delegator_id`),
  ADD KEY `delegate_id` (`delegate_id`),
  ADD KEY `policy_id` (`policy_id`),
  ADD KEY `created_by` (`created_by`);

--
-- Indexes for table `approval_notifications`
--
ALTER TABLE `approval_notifications`
  ADD PRIMARY KEY (`id`),
  ADD KEY `request_id` (`request_id`),
  ADD KEY `recipient_id` (`recipient_id`);

--
-- Indexes for table `approval_policies`
--
ALTER TABLE `approval_policies`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_policy_module` (`company_id`,`doc_module`,`is_active`),
  ADD KEY `created_by` (`created_by`);

--
-- Indexes for table `approval_requests`
--
ALTER TABLE `approval_requests`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_req_doc` (`company_id`,`doc_module`,`doc_id`),
  ADD KEY `idx_req_status` (`company_id`,`status`,`current_step`),
  ADD KEY `policy_id` (`policy_id`),
  ADD KEY `requested_by` (`requested_by`);

--
-- Indexes for table `approval_steps`
--
ALTER TABLE `approval_steps`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_step` (`policy_id`,`step_no`),
  ADD KEY `approver_role_id` (`approver_role_id`),
  ADD KEY `approver_user_id` (`approver_user_id`);

--
-- Indexes for table `audit_logs`
--
ALTER TABLE `audit_logs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_audit_table` (`company_id`,`table_name`,`record_id`),
  ADD KEY `idx_audit_user` (`company_id`,`user_id`,`created_at`);

--
-- Indexes for table `branches`
--
ALTER TABLE `branches`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_branch_code` (`company_id`,`code`);

--
-- Indexes for table `business_partners`
--
ALTER TABLE `business_partners`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_partner_code` (`company_id`,`code`),
  ADD KEY `ar_account_id` (`ar_account_id`),
  ADD KEY `ap_account_id` (`ap_account_id`);

--
-- Indexes for table `companies`
--
ALTER TABLE `companies`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Indexes for table `currencies`
--
ALTER TABLE `currencies`
  ADD PRIMARY KEY (`code`);

--
-- Indexes for table `delivery_orders`
--
ALTER TABLE `delivery_orders`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_do_no` (`company_id`,`do_no`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `so_id` (`so_id`),
  ADD KEY `customer_id` (`customer_id`),
  ADD KEY `journal_id` (`journal_id`),
  ADD KEY `created_by` (`created_by`);

--
-- Indexes for table `delivery_order_lines`
--
ALTER TABLE `delivery_order_lines`
  ADD PRIMARY KEY (`id`),
  ADD KEY `do_id` (`do_id`),
  ADD KEY `so_line_id` (`so_line_id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `warehouse_id` (`warehouse_id`),
  ADD KEY `location_id` (`location_id`),
  ADD KEY `uom_id` (`uom_id`);

--
-- Indexes for table `departments`
--
ALTER TABLE `departments`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_dept_code` (`company_id`,`code`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `parent_id` (`parent_id`);

--
-- Indexes for table `document_sequences`
--
ALTER TABLE `document_sequences`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_seq` (`company_id`,`branch_id`,`doc_type`),
  ADD KEY `branch_id` (`branch_id`);

--
-- Indexes for table `document_stamps`
--
ALTER TABLE `document_stamps`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `stamp_uuid` (`stamp_uuid`),
  ADD KEY `idx_stamp_doc` (`company_id`,`doc_module`,`doc_id`),
  ADD KEY `idx_stamp_uuid` (`stamp_uuid`),
  ADD KEY `request_id` (`request_id`),
  ADD KEY `printed_by` (`printed_by`),
  ADD KEY `voided_by` (`voided_by`);

--
-- Indexes for table `exchange_rates`
--
ALTER TABLE `exchange_rates`
  ADD PRIMARY KEY (`id`),
  ADD KEY `company_id` (`company_id`),
  ADD KEY `from_currency` (`from_currency`),
  ADD KEY `to_currency` (`to_currency`);

--
-- Indexes for table `fiscal_periods`
--
ALTER TABLE `fiscal_periods`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_period` (`company_id`,`year`,`period`),
  ADD KEY `closed_by` (`closed_by`);

--
-- Indexes for table `goods_receipts`
--
ALTER TABLE `goods_receipts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_grn_no` (`company_id`,`grn_no`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `po_id` (`po_id`),
  ADD KEY `vendor_id` (`vendor_id`),
  ADD KEY `journal_id` (`journal_id`),
  ADD KEY `created_by` (`created_by`);

--
-- Indexes for table `goods_receipt_lines`
--
ALTER TABLE `goods_receipt_lines`
  ADD PRIMARY KEY (`id`),
  ADD KEY `grn_id` (`grn_id`),
  ADD KEY `po_line_id` (`po_line_id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `warehouse_id` (`warehouse_id`),
  ADD KEY `location_id` (`location_id`),
  ADD KEY `uom_id` (`uom_id`);

--
-- Indexes for table `inventory_balances`
--
ALTER TABLE `inventory_balances`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_inv_balance` (`company_id`,`warehouse_id`,`location_id`,`product_id`),
  ADD KEY `warehouse_id` (`warehouse_id`),
  ADD KEY `location_id` (`location_id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `inventory_transactions`
--
ALTER TABLE `inventory_transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_inv_txn_product` (`company_id`,`product_id`,`txn_date`),
  ADD KEY `warehouse_id` (`warehouse_id`),
  ADD KEY `location_id` (`location_id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `journal_id` (`journal_id`),
  ADD KEY `created_by` (`created_by`);

--
-- Indexes for table `invoices`
--
ALTER TABLE `invoices`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_invoice_no` (`company_id`,`invoice_no`),
  ADD UNIQUE KEY `uq_tax_invoice_no` (`company_id`,`tax_invoice_no`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `partner_id` (`partner_id`),
  ADD KEY `journal_id` (`journal_id`),
  ADD KEY `created_by` (`created_by`),
  ADD KEY `fk_invoice_tax` (`tax_invoice_id`);

--
-- Indexes for table `invoice_lines`
--
ALTER TABLE `invoice_lines`
  ADD PRIMARY KEY (`id`),
  ADD KEY `invoice_id` (`invoice_id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `uom_id` (`uom_id`),
  ADD KEY `account_id` (`account_id`);

--
-- Indexes for table `journals`
--
ALTER TABLE `journals`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_journal_no` (`company_id`,`journal_no`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `fiscal_period_id` (`fiscal_period_id`),
  ADD KEY `created_by` (`created_by`),
  ADD KEY `posted_by` (`posted_by`);

--
-- Indexes for table `journal_lines`
--
ALTER TABLE `journal_lines`
  ADD PRIMARY KEY (`id`),
  ADD KEY `journal_id` (`journal_id`),
  ADD KEY `account_id` (`account_id`),
  ADD KEY `department_id` (`department_id`);

--
-- Indexes for table `payments`
--
ALTER TABLE `payments`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_payment_no` (`company_id`,`payment_no`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `partner_id` (`partner_id`),
  ADD KEY `journal_id` (`journal_id`),
  ADD KEY `created_by` (`created_by`),
  ADD KEY `fk_pay_wht_txn` (`wht_transaction_id`),
  ADD KEY `fk_pay_wht_cert` (`wht_certificate_id`);

--
-- Indexes for table `payment_allocations`
--
ALTER TABLE `payment_allocations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `payment_id` (`payment_id`),
  ADD KEY `invoice_id` (`invoice_id`);

--
-- Indexes for table `permissions`
--
ALTER TABLE `permissions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_perm` (`module`,`action`);

--
-- Indexes for table `po_billing_tracking`
--
ALTER TABLE `po_billing_tracking`
  ADD PRIMARY KEY (`id`),
  ADD KEY `company_id` (`company_id`),
  ADD KEY `po_id` (`po_id`),
  ADD KEY `po_line_id` (`po_line_id`),
  ADD KEY `invoice_id` (`invoice_id`),
  ADD KEY `invoice_line_id` (`invoice_line_id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_product_code` (`company_id`,`code`),
  ADD KEY `category_id` (`category_id`),
  ADD KEY `base_uom_id` (`base_uom_id`),
  ADD KEY `inventory_account_id` (`inventory_account_id`),
  ADD KEY `cogs_account_id` (`cogs_account_id`),
  ADD KEY `sales_account_id` (`sales_account_id`),
  ADD KEY `purchase_account_id` (`purchase_account_id`);

--
-- Indexes for table `product_categories`
--
ALTER TABLE `product_categories`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_cat_code` (`company_id`,`code`),
  ADD KEY `parent_id` (`parent_id`);

--
-- Indexes for table `purchase_orders`
--
ALTER TABLE `purchase_orders`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_po_no` (`company_id`,`po_no`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `department_id` (`department_id`),
  ADD KEY `vendor_id` (`vendor_id`),
  ADD KEY `approved_by` (`approved_by`),
  ADD KEY `created_by` (`created_by`);

--
-- Indexes for table `purchase_order_lines`
--
ALTER TABLE `purchase_order_lines`
  ADD PRIMARY KEY (`id`),
  ADD KEY `po_id` (`po_id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `uom_id` (`uom_id`),
  ADD KEY `warehouse_id` (`warehouse_id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_role_code` (`company_id`,`code`);

--
-- Indexes for table `role_permissions`
--
ALTER TABLE `role_permissions`
  ADD PRIMARY KEY (`role_id`,`permission_id`),
  ADD KEY `permission_id` (`permission_id`);

--
-- Indexes for table `sales_orders`
--
ALTER TABLE `sales_orders`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_so_no` (`company_id`,`so_no`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `department_id` (`department_id`),
  ADD KEY `customer_id` (`customer_id`),
  ADD KEY `salesperson_id` (`salesperson_id`),
  ADD KEY `approved_by` (`approved_by`),
  ADD KEY `created_by` (`created_by`);

--
-- Indexes for table `sales_order_lines`
--
ALTER TABLE `sales_order_lines`
  ADD PRIMARY KEY (`id`),
  ADD KEY `so_id` (`so_id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `uom_id` (`uom_id`),
  ADD KEY `warehouse_id` (`warehouse_id`);

--
-- Indexes for table `so_billing_tracking`
--
ALTER TABLE `so_billing_tracking`
  ADD PRIMARY KEY (`id`),
  ADD KEY `company_id` (`company_id`),
  ADD KEY `so_id` (`so_id`),
  ADD KEY `so_line_id` (`so_line_id`),
  ADD KEY `invoice_id` (`invoice_id`),
  ADD KEY `invoice_line_id` (`invoice_line_id`);

--
-- Indexes for table `tax_invoice_queue`
--
ALTER TABLE `tax_invoice_queue`
  ADD PRIMARY KEY (`id`),
  ADD KEY `company_id` (`company_id`),
  ADD KEY `billing_note_id` (`billing_note_id`),
  ADD KEY `payment_id` (`payment_id`),
  ADD KEY `tax_invoice_id` (`tax_invoice_id`);

--
-- Indexes for table `units_of_measure`
--
ALTER TABLE `units_of_measure`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_uom_code` (`company_id`,`code`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `user_company_access`
--
ALTER TABLE `user_company_access`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_user_access` (`user_id`,`company_id`,`branch_id`,`department_id`,`role_id`),
  ADD KEY `company_id` (`company_id`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `department_id` (`department_id`),
  ADD KEY `role_id` (`role_id`);

--
-- Indexes for table `user_signatures`
--
ALTER TABLE `user_signatures`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`);

--
-- Indexes for table `vat_rates`
--
ALTER TABLE `vat_rates`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`),
  ADD KEY `output_vat_account_id` (`output_vat_account_id`),
  ADD KEY `input_vat_account_id` (`input_vat_account_id`);

--
-- Indexes for table `vat_report_periods`
--
ALTER TABLE `vat_report_periods`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_vat_report` (`company_id`,`branch_id`,`report_year`,`report_month`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `submitted_by` (`submitted_by`);

--
-- Indexes for table `warehouses`
--
ALTER TABLE `warehouses`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_warehouse_code` (`company_id`,`code`),
  ADD KEY `branch_id` (`branch_id`);

--
-- Indexes for table `warehouse_locations`
--
ALTER TABLE `warehouse_locations`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_loc_code` (`warehouse_id`,`code`);

--
-- Indexes for table `wht_certificates`
--
ALTER TABLE `wht_certificates`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_cert_no` (`company_id`,`cert_no`),
  ADD UNIQUE KEY `stamp_uuid` (`stamp_uuid`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `partner_id` (`partner_id`),
  ADD KEY `issued_by` (`issued_by`),
  ADD KEY `created_by` (`created_by`);

--
-- Indexes for table `wht_report_periods`
--
ALTER TABLE `wht_report_periods`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_wht_report` (`company_id`,`branch_id`,`form_type`,`report_year`,`report_month`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `submitted_by` (`submitted_by`);

--
-- Indexes for table `wht_transactions`
--
ALTER TABLE `wht_transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_wht_period` (`company_id`,`fiscal_period_id`,`wht_direction`),
  ADD KEY `idx_wht_partner` (`company_id`,`partner_id`),
  ADD KEY `idx_wht_remit` (`company_id`,`is_remitted`,`payment_date`),
  ADD KEY `branch_id` (`branch_id`),
  ADD KEY `fiscal_period_id` (`fiscal_period_id`),
  ADD KEY `partner_id` (`partner_id`),
  ADD KEY `wht_type_id` (`wht_type_id`),
  ADD KEY `journal_id` (`journal_id`),
  ADD KEY `wht_account_id` (`wht_account_id`),
  ADD KEY `created_by` (`created_by`),
  ADD KEY `fk_wht_cert` (`certificate_id`),
  ADD KEY `fk_wht_report` (`report_period_id`);

--
-- Indexes for table `wht_types`
--
ALTER TABLE `wht_types`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `accounts`
--
ALTER TABLE `accounts`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=242;

--
-- AUTO_INCREMENT for table `account_types`
--
ALTER TABLE `account_types`
  MODIFY `id` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `approval_actions`
--
ALTER TABLE `approval_actions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `approval_delegates`
--
ALTER TABLE `approval_delegates`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `approval_notifications`
--
ALTER TABLE `approval_notifications`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `approval_policies`
--
ALTER TABLE `approval_policies`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `approval_requests`
--
ALTER TABLE `approval_requests`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `approval_steps`
--
ALTER TABLE `approval_steps`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `audit_logs`
--
ALTER TABLE `audit_logs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `branches`
--
ALTER TABLE `branches`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `business_partners`
--
ALTER TABLE `business_partners`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `companies`
--
ALTER TABLE `companies`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `delivery_orders`
--
ALTER TABLE `delivery_orders`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `delivery_order_lines`
--
ALTER TABLE `delivery_order_lines`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `departments`
--
ALTER TABLE `departments`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `document_sequences`
--
ALTER TABLE `document_sequences`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `document_stamps`
--
ALTER TABLE `document_stamps`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `exchange_rates`
--
ALTER TABLE `exchange_rates`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `fiscal_periods`
--
ALTER TABLE `fiscal_periods`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `goods_receipts`
--
ALTER TABLE `goods_receipts`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `goods_receipt_lines`
--
ALTER TABLE `goods_receipt_lines`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `inventory_balances`
--
ALTER TABLE `inventory_balances`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `inventory_transactions`
--
ALTER TABLE `inventory_transactions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `invoices`
--
ALTER TABLE `invoices`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `invoice_lines`
--
ALTER TABLE `invoice_lines`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `journals`
--
ALTER TABLE `journals`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `journal_lines`
--
ALTER TABLE `journal_lines`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `payments`
--
ALTER TABLE `payments`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `payment_allocations`
--
ALTER TABLE `payment_allocations`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `permissions`
--
ALTER TABLE `permissions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT for table `po_billing_tracking`
--
ALTER TABLE `po_billing_tracking`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `product_categories`
--
ALTER TABLE `product_categories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `purchase_orders`
--
ALTER TABLE `purchase_orders`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `purchase_order_lines`
--
ALTER TABLE `purchase_order_lines`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `sales_orders`
--
ALTER TABLE `sales_orders`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `sales_order_lines`
--
ALTER TABLE `sales_order_lines`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `so_billing_tracking`
--
ALTER TABLE `so_billing_tracking`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `tax_invoice_queue`
--
ALTER TABLE `tax_invoice_queue`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `units_of_measure`
--
ALTER TABLE `units_of_measure`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `user_company_access`
--
ALTER TABLE `user_company_access`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `user_signatures`
--
ALTER TABLE `user_signatures`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `vat_rates`
--
ALTER TABLE `vat_rates`
  MODIFY `id` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `vat_report_periods`
--
ALTER TABLE `vat_report_periods`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `warehouses`
--
ALTER TABLE `warehouses`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `warehouse_locations`
--
ALTER TABLE `warehouse_locations`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `wht_certificates`
--
ALTER TABLE `wht_certificates`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `wht_report_periods`
--
ALTER TABLE `wht_report_periods`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `wht_transactions`
--
ALTER TABLE `wht_transactions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `wht_types`
--
ALTER TABLE `wht_types`
  MODIFY `id` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `accounts`
--
ALTER TABLE `accounts`
  ADD CONSTRAINT `accounts_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `accounts_ibfk_2` FOREIGN KEY (`account_type_id`) REFERENCES `account_types` (`id`),
  ADD CONSTRAINT `accounts_ibfk_3` FOREIGN KEY (`parent_id`) REFERENCES `accounts` (`id`);

--
-- Constraints for table `approval_actions`
--
ALTER TABLE `approval_actions`
  ADD CONSTRAINT `approval_actions_ibfk_1` FOREIGN KEY (`request_id`) REFERENCES `approval_requests` (`id`),
  ADD CONSTRAINT `approval_actions_ibfk_2` FOREIGN KEY (`actor_user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `approval_actions_ibfk_3` FOREIGN KEY (`delegated_from`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `approval_actions_ibfk_4` FOREIGN KEY (`signature_key_id`) REFERENCES `user_signatures` (`id`);

--
-- Constraints for table `approval_delegates`
--
ALTER TABLE `approval_delegates`
  ADD CONSTRAINT `approval_delegates_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `approval_delegates_ibfk_2` FOREIGN KEY (`delegator_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `approval_delegates_ibfk_3` FOREIGN KEY (`delegate_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `approval_delegates_ibfk_4` FOREIGN KEY (`policy_id`) REFERENCES `approval_policies` (`id`),
  ADD CONSTRAINT `approval_delegates_ibfk_5` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `approval_notifications`
--
ALTER TABLE `approval_notifications`
  ADD CONSTRAINT `approval_notifications_ibfk_1` FOREIGN KEY (`request_id`) REFERENCES `approval_requests` (`id`),
  ADD CONSTRAINT `approval_notifications_ibfk_2` FOREIGN KEY (`recipient_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `approval_policies`
--
ALTER TABLE `approval_policies`
  ADD CONSTRAINT `approval_policies_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `approval_policies_ibfk_2` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `approval_requests`
--
ALTER TABLE `approval_requests`
  ADD CONSTRAINT `approval_requests_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `approval_requests_ibfk_2` FOREIGN KEY (`policy_id`) REFERENCES `approval_policies` (`id`),
  ADD CONSTRAINT `approval_requests_ibfk_3` FOREIGN KEY (`requested_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `approval_steps`
--
ALTER TABLE `approval_steps`
  ADD CONSTRAINT `approval_steps_ibfk_1` FOREIGN KEY (`policy_id`) REFERENCES `approval_policies` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `approval_steps_ibfk_2` FOREIGN KEY (`approver_role_id`) REFERENCES `roles` (`id`),
  ADD CONSTRAINT `approval_steps_ibfk_3` FOREIGN KEY (`approver_user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `audit_logs`
--
ALTER TABLE `audit_logs`
  ADD CONSTRAINT `audit_logs_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`);

--
-- Constraints for table `branches`
--
ALTER TABLE `branches`
  ADD CONSTRAINT `branches_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`);

--
-- Constraints for table `business_partners`
--
ALTER TABLE `business_partners`
  ADD CONSTRAINT `business_partners_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `business_partners_ibfk_2` FOREIGN KEY (`ar_account_id`) REFERENCES `accounts` (`id`),
  ADD CONSTRAINT `business_partners_ibfk_3` FOREIGN KEY (`ap_account_id`) REFERENCES `accounts` (`id`);

--
-- Constraints for table `delivery_orders`
--
ALTER TABLE `delivery_orders`
  ADD CONSTRAINT `delivery_orders_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `delivery_orders_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `delivery_orders_ibfk_3` FOREIGN KEY (`so_id`) REFERENCES `sales_orders` (`id`),
  ADD CONSTRAINT `delivery_orders_ibfk_4` FOREIGN KEY (`customer_id`) REFERENCES `business_partners` (`id`),
  ADD CONSTRAINT `delivery_orders_ibfk_5` FOREIGN KEY (`journal_id`) REFERENCES `journals` (`id`),
  ADD CONSTRAINT `delivery_orders_ibfk_6` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `delivery_order_lines`
--
ALTER TABLE `delivery_order_lines`
  ADD CONSTRAINT `delivery_order_lines_ibfk_1` FOREIGN KEY (`do_id`) REFERENCES `delivery_orders` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `delivery_order_lines_ibfk_2` FOREIGN KEY (`so_line_id`) REFERENCES `sales_order_lines` (`id`),
  ADD CONSTRAINT `delivery_order_lines_ibfk_3` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `delivery_order_lines_ibfk_4` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`),
  ADD CONSTRAINT `delivery_order_lines_ibfk_5` FOREIGN KEY (`location_id`) REFERENCES `warehouse_locations` (`id`),
  ADD CONSTRAINT `delivery_order_lines_ibfk_6` FOREIGN KEY (`uom_id`) REFERENCES `units_of_measure` (`id`);

--
-- Constraints for table `departments`
--
ALTER TABLE `departments`
  ADD CONSTRAINT `departments_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `departments_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `departments_ibfk_3` FOREIGN KEY (`parent_id`) REFERENCES `departments` (`id`);

--
-- Constraints for table `document_sequences`
--
ALTER TABLE `document_sequences`
  ADD CONSTRAINT `document_sequences_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `document_sequences_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`);

--
-- Constraints for table `document_stamps`
--
ALTER TABLE `document_stamps`
  ADD CONSTRAINT `document_stamps_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `document_stamps_ibfk_2` FOREIGN KEY (`request_id`) REFERENCES `approval_requests` (`id`),
  ADD CONSTRAINT `document_stamps_ibfk_3` FOREIGN KEY (`printed_by`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `document_stamps_ibfk_4` FOREIGN KEY (`voided_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `exchange_rates`
--
ALTER TABLE `exchange_rates`
  ADD CONSTRAINT `exchange_rates_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `exchange_rates_ibfk_2` FOREIGN KEY (`from_currency`) REFERENCES `currencies` (`code`),
  ADD CONSTRAINT `exchange_rates_ibfk_3` FOREIGN KEY (`to_currency`) REFERENCES `currencies` (`code`);

--
-- Constraints for table `fiscal_periods`
--
ALTER TABLE `fiscal_periods`
  ADD CONSTRAINT `fiscal_periods_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `fiscal_periods_ibfk_2` FOREIGN KEY (`closed_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `goods_receipts`
--
ALTER TABLE `goods_receipts`
  ADD CONSTRAINT `goods_receipts_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `goods_receipts_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `goods_receipts_ibfk_3` FOREIGN KEY (`po_id`) REFERENCES `purchase_orders` (`id`),
  ADD CONSTRAINT `goods_receipts_ibfk_4` FOREIGN KEY (`vendor_id`) REFERENCES `business_partners` (`id`),
  ADD CONSTRAINT `goods_receipts_ibfk_5` FOREIGN KEY (`journal_id`) REFERENCES `journals` (`id`),
  ADD CONSTRAINT `goods_receipts_ibfk_6` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `goods_receipt_lines`
--
ALTER TABLE `goods_receipt_lines`
  ADD CONSTRAINT `goods_receipt_lines_ibfk_1` FOREIGN KEY (`grn_id`) REFERENCES `goods_receipts` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `goods_receipt_lines_ibfk_2` FOREIGN KEY (`po_line_id`) REFERENCES `purchase_order_lines` (`id`),
  ADD CONSTRAINT `goods_receipt_lines_ibfk_3` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `goods_receipt_lines_ibfk_4` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`),
  ADD CONSTRAINT `goods_receipt_lines_ibfk_5` FOREIGN KEY (`location_id`) REFERENCES `warehouse_locations` (`id`),
  ADD CONSTRAINT `goods_receipt_lines_ibfk_6` FOREIGN KEY (`uom_id`) REFERENCES `units_of_measure` (`id`);

--
-- Constraints for table `inventory_balances`
--
ALTER TABLE `inventory_balances`
  ADD CONSTRAINT `inventory_balances_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `inventory_balances_ibfk_2` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`),
  ADD CONSTRAINT `inventory_balances_ibfk_3` FOREIGN KEY (`location_id`) REFERENCES `warehouse_locations` (`id`),
  ADD CONSTRAINT `inventory_balances_ibfk_4` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

--
-- Constraints for table `inventory_transactions`
--
ALTER TABLE `inventory_transactions`
  ADD CONSTRAINT `inventory_transactions_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `inventory_transactions_ibfk_2` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`),
  ADD CONSTRAINT `inventory_transactions_ibfk_3` FOREIGN KEY (`location_id`) REFERENCES `warehouse_locations` (`id`),
  ADD CONSTRAINT `inventory_transactions_ibfk_4` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `inventory_transactions_ibfk_5` FOREIGN KEY (`journal_id`) REFERENCES `journals` (`id`),
  ADD CONSTRAINT `inventory_transactions_ibfk_6` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `invoices`
--
ALTER TABLE `invoices`
  ADD CONSTRAINT `fk_invoice_tax` FOREIGN KEY (`tax_invoice_id`) REFERENCES `invoices` (`id`),
  ADD CONSTRAINT `invoices_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `invoices_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `invoices_ibfk_3` FOREIGN KEY (`partner_id`) REFERENCES `business_partners` (`id`),
  ADD CONSTRAINT `invoices_ibfk_4` FOREIGN KEY (`journal_id`) REFERENCES `journals` (`id`),
  ADD CONSTRAINT `invoices_ibfk_5` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `invoice_lines`
--
ALTER TABLE `invoice_lines`
  ADD CONSTRAINT `invoice_lines_ibfk_1` FOREIGN KEY (`invoice_id`) REFERENCES `invoices` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `invoice_lines_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `invoice_lines_ibfk_3` FOREIGN KEY (`uom_id`) REFERENCES `units_of_measure` (`id`),
  ADD CONSTRAINT `invoice_lines_ibfk_4` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

--
-- Constraints for table `journals`
--
ALTER TABLE `journals`
  ADD CONSTRAINT `journals_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `journals_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `journals_ibfk_3` FOREIGN KEY (`fiscal_period_id`) REFERENCES `fiscal_periods` (`id`),
  ADD CONSTRAINT `journals_ibfk_4` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `journals_ibfk_5` FOREIGN KEY (`posted_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `journal_lines`
--
ALTER TABLE `journal_lines`
  ADD CONSTRAINT `journal_lines_ibfk_1` FOREIGN KEY (`journal_id`) REFERENCES `journals` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `journal_lines_ibfk_2` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`),
  ADD CONSTRAINT `journal_lines_ibfk_3` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`);

--
-- Constraints for table `payments`
--
ALTER TABLE `payments`
  ADD CONSTRAINT `fk_pay_wht_cert` FOREIGN KEY (`wht_certificate_id`) REFERENCES `wht_certificates` (`id`),
  ADD CONSTRAINT `fk_pay_wht_txn` FOREIGN KEY (`wht_transaction_id`) REFERENCES `wht_transactions` (`id`),
  ADD CONSTRAINT `payments_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `payments_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `payments_ibfk_3` FOREIGN KEY (`partner_id`) REFERENCES `business_partners` (`id`),
  ADD CONSTRAINT `payments_ibfk_4` FOREIGN KEY (`journal_id`) REFERENCES `journals` (`id`),
  ADD CONSTRAINT `payments_ibfk_5` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `payment_allocations`
--
ALTER TABLE `payment_allocations`
  ADD CONSTRAINT `payment_allocations_ibfk_1` FOREIGN KEY (`payment_id`) REFERENCES `payments` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `payment_allocations_ibfk_2` FOREIGN KEY (`invoice_id`) REFERENCES `invoices` (`id`);

--
-- Constraints for table `po_billing_tracking`
--
ALTER TABLE `po_billing_tracking`
  ADD CONSTRAINT `po_billing_tracking_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `po_billing_tracking_ibfk_2` FOREIGN KEY (`po_id`) REFERENCES `purchase_orders` (`id`),
  ADD CONSTRAINT `po_billing_tracking_ibfk_3` FOREIGN KEY (`po_line_id`) REFERENCES `purchase_order_lines` (`id`),
  ADD CONSTRAINT `po_billing_tracking_ibfk_4` FOREIGN KEY (`invoice_id`) REFERENCES `invoices` (`id`),
  ADD CONSTRAINT `po_billing_tracking_ibfk_5` FOREIGN KEY (`invoice_line_id`) REFERENCES `invoice_lines` (`id`);

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `product_categories` (`id`),
  ADD CONSTRAINT `products_ibfk_3` FOREIGN KEY (`base_uom_id`) REFERENCES `units_of_measure` (`id`),
  ADD CONSTRAINT `products_ibfk_4` FOREIGN KEY (`inventory_account_id`) REFERENCES `accounts` (`id`),
  ADD CONSTRAINT `products_ibfk_5` FOREIGN KEY (`cogs_account_id`) REFERENCES `accounts` (`id`),
  ADD CONSTRAINT `products_ibfk_6` FOREIGN KEY (`sales_account_id`) REFERENCES `accounts` (`id`),
  ADD CONSTRAINT `products_ibfk_7` FOREIGN KEY (`purchase_account_id`) REFERENCES `accounts` (`id`);

--
-- Constraints for table `product_categories`
--
ALTER TABLE `product_categories`
  ADD CONSTRAINT `product_categories_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `product_categories_ibfk_2` FOREIGN KEY (`parent_id`) REFERENCES `product_categories` (`id`);

--
-- Constraints for table `purchase_orders`
--
ALTER TABLE `purchase_orders`
  ADD CONSTRAINT `purchase_orders_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `purchase_orders_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `purchase_orders_ibfk_3` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`),
  ADD CONSTRAINT `purchase_orders_ibfk_4` FOREIGN KEY (`vendor_id`) REFERENCES `business_partners` (`id`),
  ADD CONSTRAINT `purchase_orders_ibfk_5` FOREIGN KEY (`approved_by`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `purchase_orders_ibfk_6` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `purchase_order_lines`
--
ALTER TABLE `purchase_order_lines`
  ADD CONSTRAINT `purchase_order_lines_ibfk_1` FOREIGN KEY (`po_id`) REFERENCES `purchase_orders` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `purchase_order_lines_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `purchase_order_lines_ibfk_3` FOREIGN KEY (`uom_id`) REFERENCES `units_of_measure` (`id`),
  ADD CONSTRAINT `purchase_order_lines_ibfk_4` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`);

--
-- Constraints for table `roles`
--
ALTER TABLE `roles`
  ADD CONSTRAINT `roles_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`);

--
-- Constraints for table `role_permissions`
--
ALTER TABLE `role_permissions`
  ADD CONSTRAINT `role_permissions_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `role_permissions_ibfk_2` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `sales_orders`
--
ALTER TABLE `sales_orders`
  ADD CONSTRAINT `sales_orders_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `sales_orders_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `sales_orders_ibfk_3` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`),
  ADD CONSTRAINT `sales_orders_ibfk_4` FOREIGN KEY (`customer_id`) REFERENCES `business_partners` (`id`),
  ADD CONSTRAINT `sales_orders_ibfk_5` FOREIGN KEY (`salesperson_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `sales_orders_ibfk_6` FOREIGN KEY (`approved_by`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `sales_orders_ibfk_7` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `sales_order_lines`
--
ALTER TABLE `sales_order_lines`
  ADD CONSTRAINT `sales_order_lines_ibfk_1` FOREIGN KEY (`so_id`) REFERENCES `sales_orders` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `sales_order_lines_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `sales_order_lines_ibfk_3` FOREIGN KEY (`uom_id`) REFERENCES `units_of_measure` (`id`),
  ADD CONSTRAINT `sales_order_lines_ibfk_4` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`);

--
-- Constraints for table `so_billing_tracking`
--
ALTER TABLE `so_billing_tracking`
  ADD CONSTRAINT `so_billing_tracking_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `so_billing_tracking_ibfk_2` FOREIGN KEY (`so_id`) REFERENCES `sales_orders` (`id`),
  ADD CONSTRAINT `so_billing_tracking_ibfk_3` FOREIGN KEY (`so_line_id`) REFERENCES `sales_order_lines` (`id`),
  ADD CONSTRAINT `so_billing_tracking_ibfk_4` FOREIGN KEY (`invoice_id`) REFERENCES `invoices` (`id`),
  ADD CONSTRAINT `so_billing_tracking_ibfk_5` FOREIGN KEY (`invoice_line_id`) REFERENCES `invoice_lines` (`id`);

--
-- Constraints for table `tax_invoice_queue`
--
ALTER TABLE `tax_invoice_queue`
  ADD CONSTRAINT `tax_invoice_queue_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `tax_invoice_queue_ibfk_2` FOREIGN KEY (`billing_note_id`) REFERENCES `invoices` (`id`),
  ADD CONSTRAINT `tax_invoice_queue_ibfk_3` FOREIGN KEY (`payment_id`) REFERENCES `payments` (`id`),
  ADD CONSTRAINT `tax_invoice_queue_ibfk_4` FOREIGN KEY (`tax_invoice_id`) REFERENCES `invoices` (`id`);

--
-- Constraints for table `units_of_measure`
--
ALTER TABLE `units_of_measure`
  ADD CONSTRAINT `units_of_measure_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`);

--
-- Constraints for table `user_company_access`
--
ALTER TABLE `user_company_access`
  ADD CONSTRAINT `user_company_access_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `user_company_access_ibfk_2` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `user_company_access_ibfk_3` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `user_company_access_ibfk_4` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`),
  ADD CONSTRAINT `user_company_access_ibfk_5` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

--
-- Constraints for table `user_signatures`
--
ALTER TABLE `user_signatures`
  ADD CONSTRAINT `user_signatures_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `vat_rates`
--
ALTER TABLE `vat_rates`
  ADD CONSTRAINT `vat_rates_ibfk_1` FOREIGN KEY (`output_vat_account_id`) REFERENCES `accounts` (`id`),
  ADD CONSTRAINT `vat_rates_ibfk_2` FOREIGN KEY (`input_vat_account_id`) REFERENCES `accounts` (`id`);

--
-- Constraints for table `vat_report_periods`
--
ALTER TABLE `vat_report_periods`
  ADD CONSTRAINT `vat_report_periods_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `vat_report_periods_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `vat_report_periods_ibfk_3` FOREIGN KEY (`submitted_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `warehouses`
--
ALTER TABLE `warehouses`
  ADD CONSTRAINT `warehouses_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `warehouses_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`);

--
-- Constraints for table `warehouse_locations`
--
ALTER TABLE `warehouse_locations`
  ADD CONSTRAINT `warehouse_locations_ibfk_1` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`);

--
-- Constraints for table `wht_certificates`
--
ALTER TABLE `wht_certificates`
  ADD CONSTRAINT `wht_certificates_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `wht_certificates_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `wht_certificates_ibfk_3` FOREIGN KEY (`partner_id`) REFERENCES `business_partners` (`id`),
  ADD CONSTRAINT `wht_certificates_ibfk_4` FOREIGN KEY (`issued_by`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `wht_certificates_ibfk_5` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `wht_report_periods`
--
ALTER TABLE `wht_report_periods`
  ADD CONSTRAINT `wht_report_periods_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `wht_report_periods_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `wht_report_periods_ibfk_3` FOREIGN KEY (`submitted_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `wht_transactions`
--
ALTER TABLE `wht_transactions`
  ADD CONSTRAINT `fk_wht_cert` FOREIGN KEY (`certificate_id`) REFERENCES `wht_certificates` (`id`),
  ADD CONSTRAINT `fk_wht_report` FOREIGN KEY (`report_period_id`) REFERENCES `wht_report_periods` (`id`),
  ADD CONSTRAINT `wht_transactions_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
  ADD CONSTRAINT `wht_transactions_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`),
  ADD CONSTRAINT `wht_transactions_ibfk_3` FOREIGN KEY (`fiscal_period_id`) REFERENCES `fiscal_periods` (`id`),
  ADD CONSTRAINT `wht_transactions_ibfk_4` FOREIGN KEY (`partner_id`) REFERENCES `business_partners` (`id`),
  ADD CONSTRAINT `wht_transactions_ibfk_5` FOREIGN KEY (`wht_type_id`) REFERENCES `wht_types` (`id`),
  ADD CONSTRAINT `wht_transactions_ibfk_6` FOREIGN KEY (`journal_id`) REFERENCES `journals` (`id`),
  ADD CONSTRAINT `wht_transactions_ibfk_7` FOREIGN KEY (`wht_account_id`) REFERENCES `accounts` (`id`),
  ADD CONSTRAINT `wht_transactions_ibfk_8` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
